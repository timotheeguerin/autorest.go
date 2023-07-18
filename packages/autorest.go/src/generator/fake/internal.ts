/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@autorest/extension-base';
import { CodeModel } from '@autorest/codemodel';
import { values } from '@azure-tools/linq';
import { contentPreamble } from '../helpers';
import { ImportManager } from '../imports';
import { isLROOperation, isPageableOperation } from '../../common/helpers';

export async function generateServerInternal(session: Session<CodeModel>): Promise<string> {
  if (session.model.operationGroups.length === 0) {
    return '';
  }
  const text = await contentPreamble(session, 'fake');
  const imports = new ImportManager();
  imports.add('io');
  imports.add('net/http');
  imports.add('reflect');
  let body = content;
  // only generate the tracker content if required
  let needsTracker = false;
  for (const group of values(session.model.operationGroups)) {
    for (const op of values(group.operations)) {
      if (isLROOperation(op) || isPageableOperation(op)) {
        needsTracker = true;
        break;
      }
    }
    if (needsTracker) {
      break;
    }
  }
  if (needsTracker) {
    imports.add('regexp');
    imports.add('strings');
    imports.add('sync');
    body += tracker;
  }
  return text + imports.text() + body;
}

const content = `
type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}

func getHeaderValue(h http.Header, k string) string {
	v := h[k]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func parseOptional[T any](v string, parse func(v string) (T, error)) (*T, error) {
	if v == "" {
		return nil, nil
	}
	t, err := parse(v)
	if err != nil {
		return nil, err
	}
	return &t, err
}

func parseWithCast[T any](v string, parse func(v string) (T, error)) (T, error) {
	t, err := parse(v)
	if err != nil {
		return *new(T), err
	}
	return t, err
}

func readRequestBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body.Close()
	return body, nil
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
`;

const tracker = `
func newTracker[T any]() *tracker[T] {
	return &tracker[T]{
		items: map[string]*T{},
	}
}

type tracker[T any] struct {
	items map[string]*T
	mu sync.Mutex
}

func (p *tracker[T]) key(req *http.Request) string {
	path := req.URL.Path
	if match, _ := regexp.Match(\`/page_\\d+$\`, []byte(path)); match {
		path = path[:strings.LastIndex(path, "/")]
	} else if strings.HasSuffix(path, "/get/fake/status") {
		path = path[:len(path)-16]
	}
	return path
}

func (p *tracker[T]) get(req *http.Request) *T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if item, ok := p.items[p.key(req)]; ok {
		return item
	}
	return nil
}

func (p *tracker[T]) add(req *http.Request, item *T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.items[p.key(req)] = item
}

func (p *tracker[T]) remove(req *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.items, p.key(req))
}
`;
