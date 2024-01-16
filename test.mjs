import { resolve } from 'node:path'
import { realpath, writeFile } from 'node:fs/promises'
import { realpath as realpathCallback, realpathSync } from 'node:fs'

const p = `./common/temp/node_modules/.pnpm/@azure-tools+cadl-ranch-specs@0.26.2_@azure-tools+cadl-ranch-expect@0.9.0_@azure-tools+typesp_6mmb4jnmojqwtbe6z5fh4mcmqi/node_modules/@azure-tools/cadl-ranch-specs/http/client/structure/common/service.tsp`

await writeFile(p, '')

console.log('length: ', resolve(p).length)

let hasError = false
try {
  await realpath(p)
} catch (e) {
  console.log('error with fs.promises.realpath', e)
  hasError = true;
}

try {
  await new Promise((res, rej) => {
    realpathCallback(p, (err, resolved) => {
      if (err) {
        rej(err)
        return
      }
      res(resolved)
    })
  })
} catch (e) {
  console.log('error with fs.realpath', e)
  hasError = true;

}

try {
  await new Promise((res, rej) => {
    realpathCallback.native(p, (err, resolved) => {
      if (err) {
        rej(err)
        return
      }
      res(resolved)
    })
  })
} catch (e) {
  console.log('error with fs.realpath.native', e)
}

try {
  realpathSync(p)
} catch (e) {
  console.log('error with fs.realpathSync', e)
  hasError = true;

}

try {
  realpathSync.native(p)
} catch (e) {
  console.log('error with fs.realpathSync.native', e)
  hasError = true;

}

if(hasError) {
  process.exit(1);
}