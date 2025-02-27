//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

// AssetReferenceBaseClassification provides polymorphic access to related types.
// Call the interface's GetAssetReferenceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AssetReferenceBase, *DataPathAssetReference, *IDAssetReference, *OutputPathAssetReference
type AssetReferenceBaseClassification interface {
	// GetAssetReferenceBase returns the AssetReferenceBase content of the underlying type.
	GetAssetReferenceBase() *AssetReferenceBase
}

// AutoMLVerticalClassification provides polymorphic access to related types.
// Call the interface's GetAutoMLVertical() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoMLVertical, *Classification, *Forecasting, *ImageClassification, *ImageClassificationMultilabel, *ImageInstanceSegmentation,
// - *ImageObjectDetection, *Regression, *TextClassification, *TextClassificationMultilabel, *TextNer
type AutoMLVerticalClassification interface {
	// GetAutoMLVertical returns the AutoMLVertical content of the underlying type.
	GetAutoMLVertical() *AutoMLVertical
}

// ComputeClassification provides polymorphic access to related types.
// Call the interface's GetCompute() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AKS, *AmlCompute, *Compute, *ComputeInstance, *DataFactory, *DataLakeAnalytics, *Databricks, *HDInsight, *Kubernetes,
// - *SynapseSpark, *VirtualMachine
type ComputeClassification interface {
	// GetCompute returns the Compute content of the underlying type.
	GetCompute() *Compute
}

// ComputeSecretsClassification provides polymorphic access to related types.
// Call the interface's GetComputeSecrets() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AksComputeSecrets, *ComputeSecrets, *DatabricksComputeSecrets, *VirtualMachineSecrets
type ComputeSecretsClassification interface {
	// GetComputeSecrets returns the ComputeSecrets content of the underlying type.
	GetComputeSecrets() *ComputeSecrets
}

// DataVersionBaseDetailsClassification provides polymorphic access to related types.
// Call the interface's GetDataVersionBaseDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataVersionBaseDetails, *MLTableData, *URIFileDataVersion, *URIFolderDataVersion
type DataVersionBaseDetailsClassification interface {
	// GetDataVersionBaseDetails returns the DataVersionBaseDetails content of the underlying type.
	GetDataVersionBaseDetails() *DataVersionBaseDetails
}

// DatastoreCredentialsClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreCredentials() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccountKeyDatastoreCredentials, *CertificateDatastoreCredentials, *DatastoreCredentials, *KerberosKeytabCredentials,
// - *KerberosPasswordCredentials, *NoneDatastoreCredentials, *SasDatastoreCredentials, *ServicePrincipalDatastoreCredentials
type DatastoreCredentialsClassification interface {
	// GetDatastoreCredentials returns the DatastoreCredentials content of the underlying type.
	GetDatastoreCredentials() *DatastoreCredentials
}

// DatastoreDetailsClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureBlobDatastore, *AzureDataLakeGen1Datastore, *AzureDataLakeGen2Datastore, *AzureFileDatastore, *DatastoreDetails,
// - *HdfsDatastore
type DatastoreDetailsClassification interface {
	// GetDatastoreDetails returns the DatastoreDetails content of the underlying type.
	GetDatastoreDetails() *DatastoreDetails
}

// DatastoreSecretsClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreSecrets() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccountKeyDatastoreSecrets, *CertificateDatastoreSecrets, *DatastoreSecrets, *KerberosKeytabSecrets, *KerberosPasswordSecrets,
// - *SasDatastoreSecrets, *ServicePrincipalDatastoreSecrets
type DatastoreSecretsClassification interface {
	// GetDatastoreSecrets returns the DatastoreSecrets content of the underlying type.
	GetDatastoreSecrets() *DatastoreSecrets
}

// DistributionConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetDistributionConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DistributionConfiguration, *Mpi, *PyTorch, *TensorFlow
type DistributionConfigurationClassification interface {
	// GetDistributionConfiguration returns the DistributionConfiguration content of the underlying type.
	GetDistributionConfiguration() *DistributionConfiguration
}

// EarlyTerminationPolicyClassification provides polymorphic access to related types.
// Call the interface's GetEarlyTerminationPolicy() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BanditPolicy, *EarlyTerminationPolicy, *MedianStoppingPolicy, *TruncationSelectionPolicy
type EarlyTerminationPolicyClassification interface {
	// GetEarlyTerminationPolicy returns the EarlyTerminationPolicy content of the underlying type.
	GetEarlyTerminationPolicy() *EarlyTerminationPolicy
}

// ForecastHorizonClassification provides polymorphic access to related types.
// Call the interface's GetForecastHorizon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoForecastHorizon, *CustomForecastHorizon, *ForecastHorizon
type ForecastHorizonClassification interface {
	// GetForecastHorizon returns the ForecastHorizon content of the underlying type.
	GetForecastHorizon() *ForecastHorizon
}

// IdentityConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetIdentityConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmlToken, *IdentityConfiguration, *ManagedIdentity, *UserIdentity
type IdentityConfigurationClassification interface {
	// GetIdentityConfiguration returns the IdentityConfiguration content of the underlying type.
	GetIdentityConfiguration() *IdentityConfiguration
}

// JobBaseDetailsClassification provides polymorphic access to related types.
// Call the interface's GetJobBaseDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoMLJob, *CommandJob, *JobBaseDetails, *PipelineJob, *SweepJob
type JobBaseDetailsClassification interface {
	// GetJobBaseDetails returns the JobBaseDetails content of the underlying type.
	GetJobBaseDetails() *JobBaseDetails
}

// JobInputClassification provides polymorphic access to related types.
// Call the interface's GetJobInput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomModelJobInput, *JobInput, *LiteralJobInput, *MLFlowModelJobInput, *MLTableJobInput, *TritonModelJobInput, *URIFileJobInput,
// - *URIFolderJobInput
type JobInputClassification interface {
	// GetJobInput returns the JobInput content of the underlying type.
	GetJobInput() *JobInput
}

// JobLimitsClassification provides polymorphic access to related types.
// Call the interface's GetJobLimits() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CommandJobLimits, *JobLimits, *SweepJobLimits
type JobLimitsClassification interface {
	// GetJobLimits returns the JobLimits content of the underlying type.
	GetJobLimits() *JobLimits
}

// JobOutputClassification provides polymorphic access to related types.
// Call the interface's GetJobOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomModelJobOutput, *JobOutput, *MLFlowModelJobOutput, *MLTableJobOutput, *TritonModelJobOutput, *URIFileJobOutput,
// - *URIFolderJobOutput
type JobOutputClassification interface {
	// GetJobOutput returns the JobOutput content of the underlying type.
	GetJobOutput() *JobOutput
}

// NCrossValidationsClassification provides polymorphic access to related types.
// Call the interface's GetNCrossValidations() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoNCrossValidations, *CustomNCrossValidations, *NCrossValidations
type NCrossValidationsClassification interface {
	// GetNCrossValidations returns the NCrossValidations content of the underlying type.
	GetNCrossValidations() *NCrossValidations
}

// OnlineDeploymentDetailsClassification provides polymorphic access to related types.
// Call the interface's GetOnlineDeploymentDetails() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *KubernetesOnlineDeployment, *ManagedOnlineDeployment, *OnlineDeploymentDetails
type OnlineDeploymentDetailsClassification interface {
	// GetOnlineDeploymentDetails returns the OnlineDeploymentDetails content of the underlying type.
	GetOnlineDeploymentDetails() *OnlineDeploymentDetails
}

// OnlineScaleSettingsClassification provides polymorphic access to related types.
// Call the interface's GetOnlineScaleSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DefaultScaleSettings, *OnlineScaleSettings, *TargetUtilizationScaleSettings
type OnlineScaleSettingsClassification interface {
	// GetOnlineScaleSettings returns the OnlineScaleSettings content of the underlying type.
	GetOnlineScaleSettings() *OnlineScaleSettings
}

// PartialAssetReferenceBaseClassification provides polymorphic access to related types.
// Call the interface's GetPartialAssetReferenceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PartialAssetReferenceBase, *PartialDataPathAssetReference, *PartialIDAssetReference, *PartialOutputPathAssetReference
type PartialAssetReferenceBaseClassification interface {
	// GetPartialAssetReferenceBase returns the PartialAssetReferenceBase content of the underlying type.
	GetPartialAssetReferenceBase() *PartialAssetReferenceBase
}

// PartialOnlineDeploymentClassification provides polymorphic access to related types.
// Call the interface's GetPartialOnlineDeployment() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PartialKubernetesOnlineDeployment, *PartialManagedOnlineDeployment, *PartialOnlineDeployment
type PartialOnlineDeploymentClassification interface {
	// GetPartialOnlineDeployment returns the PartialOnlineDeployment content of the underlying type.
	GetPartialOnlineDeployment() *PartialOnlineDeployment
}

// SamplingAlgorithmClassification provides polymorphic access to related types.
// Call the interface's GetSamplingAlgorithm() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BayesianSamplingAlgorithm, *GridSamplingAlgorithm, *RandomSamplingAlgorithm, *SamplingAlgorithm
type SamplingAlgorithmClassification interface {
	// GetSamplingAlgorithm returns the SamplingAlgorithm content of the underlying type.
	GetSamplingAlgorithm() *SamplingAlgorithm
}

// ScheduleBaseClassification provides polymorphic access to related types.
// Call the interface's GetScheduleBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CronSchedule, *RecurrenceSchedule, *ScheduleBase
type ScheduleBaseClassification interface {
	// GetScheduleBase returns the ScheduleBase content of the underlying type.
	GetScheduleBase() *ScheduleBase
}

// SeasonalityClassification provides polymorphic access to related types.
// Call the interface's GetSeasonality() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoSeasonality, *CustomSeasonality, *Seasonality
type SeasonalityClassification interface {
	// GetSeasonality returns the Seasonality content of the underlying type.
	GetSeasonality() *Seasonality
}

// TargetLagsClassification provides polymorphic access to related types.
// Call the interface's GetTargetLags() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoTargetLags, *CustomTargetLags, *TargetLags
type TargetLagsClassification interface {
	// GetTargetLags returns the TargetLags content of the underlying type.
	GetTargetLags() *TargetLags
}

// TargetRollingWindowSizeClassification provides polymorphic access to related types.
// Call the interface's GetTargetRollingWindowSize() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoTargetRollingWindowSize, *CustomTargetRollingWindowSize, *TargetRollingWindowSize
type TargetRollingWindowSizeClassification interface {
	// GetTargetRollingWindowSize returns the TargetRollingWindowSize content of the underlying type.
	GetTargetRollingWindowSize() *TargetRollingWindowSize
}
