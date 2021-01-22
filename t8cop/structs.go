package t8cop

type Acims_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Action_orchestrator_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Actionscript_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Aix_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Api_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Apic_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Appdynamics_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Appinsights_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Arangodb_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		Apps *string `yaml:"apps,omitempty"`
		Arangodb *string `yaml:"arangodb,omitempty"`
		Dump *string `yaml:"dump,omitempty"`
	} `yaml:"persistence,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Auth_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Aws_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Awslambda_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Azure_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Baremetal_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Cassandra_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	ArgsOverrides *string `yaml:"argsOverrides,omitempty"`
	Backup struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Destination *string `yaml:"destination,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
	} `yaml:"backup,omitempty"`
	CommandOverrides *string `yaml:"commandOverrides,omitempty"`
	Config struct {
		Cluster_domain *string `yaml:"cluster_domain,omitempty"`
		Cluster_name *string `yaml:"cluster_name,omitempty"`
		Cluster_size *int64 `yaml:"cluster_size,omitempty"`
		Dc_name *string `yaml:"dc_name,omitempty"`
		Endpoint_snitch *string `yaml:"endpoint_snitch,omitempty"`
		Heap_new_size *string `yaml:"heap_new_size,omitempty"`
		Max_heap_size *string `yaml:"max_heap_size,omitempty"`
		Num_tokens *int64 `yaml:"num_tokens,omitempty"`
		Ports struct {
			Agent *string `yaml:"agent,omitempty"`
			Cql *int64 `yaml:"cql,omitempty"`
			Thrift *int64 `yaml:"thrift,omitempty"`
		} `yaml:"ports,omitempty"`
		Rack_name *string `yaml:"rack_name,omitempty"`
		Seed_size *int64 `yaml:"seed_size,omitempty"`
		Seeds *string `yaml:"seeds,omitempty"`
		Start_rpc *bool `yaml:"start_rpc,omitempty"`
	} `yaml:"config,omitempty"`
	ConfigOverrides GenericMap `yaml:"configOverrides,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	Exporter struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			Repo *string `yaml:"repo,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		JvmOpts *string `yaml:"jvmOpts,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
	} `yaml:"exporter,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	HostNetwork *bool `yaml:"hostNetwork,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets *string `yaml:"pullSecrets,omitempty"`
		Repo *string `yaml:"repo,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	LivenessProbe struct {
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		AccessMode *string `yaml:"accessMode,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClass *string `yaml:"storageClass,omitempty"`
	} `yaml:"persistence,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodDisruptionBudget GenericMap `yaml:"podDisruptionBudget,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
	PodSettings struct {
		TerminationGracePeriodSeconds *int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
	} `yaml:"podSettings,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
	} `yaml:"rbac,omitempty"`
	ReadinessProbe struct {
		Address *string `yaml:"address,omitempty"`
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	SchedulerName *string `yaml:"schedulerName,omitempty"`
	SecurityContext struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Selector GenericMap `yaml:"selector,omitempty"`
	Service struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceAccount struct {
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
	UpdateStrategy struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"updateStrategy,omitempty"`
}

type Chronograf_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	EnvFromSecret *string `yaml:"envFromSecret,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Ingress struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Hostname *string `yaml:"hostname,omitempty"`
		SecretName *string `yaml:"secretName,omitempty"`
		Tls *bool `yaml:"tls,omitempty"`
	} `yaml:"ingress,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	Oauth struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Github struct {
			Client_id *string `yaml:"client_id,omitempty"`
			Client_secret *string `yaml:"client_secret,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Gh_orgs *string `yaml:"gh_orgs,omitempty"`
		} `yaml:"github,omitempty"`
		Google struct {
			Client_id *string `yaml:"client_id,omitempty"`
			Client_secret *string `yaml:"client_secret,omitempty"`
			Domains *string `yaml:"domains,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Public_url *string `yaml:"public_url,omitempty"`
		} `yaml:"google,omitempty"`
		Heroku struct {
			Client_id *string `yaml:"client_id,omitempty"`
			Client_secret *string `yaml:"client_secret,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			He_orgs *string `yaml:"he_orgs,omitempty"`
			Heroku_orgs *string `yaml:"heroku_orgs,omitempty"`
		} `yaml:"heroku,omitempty"`
		Token_secret *string `yaml:"token_secret,omitempty"`
	} `yaml:"oauth,omitempty"`
	Persistence struct {
		AccessMode *string `yaml:"accessMode,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClass *string `yaml:"storageClass,omitempty"`
	} `yaml:"persistence,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	Service struct {
		Port *int64 `yaml:"port,omitempty"`
		Replicas *int64 `yaml:"replicas,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
}

type Cloudfoundry_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Clustermgr_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Compellent_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Consul_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Control_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Cost_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Customdata_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Datacloud_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Collector struct {
		Api struct {
			Action struct {
				Enable *string `yaml:"enable,omitempty"`
				Host *string `yaml:"host,omitempty"`
				Port *int64 `yaml:"port,omitempty"`
			} `yaml:"action,omitempty"`
			Group struct {
				Host *string `yaml:"host,omitempty"`
				Port *int64 `yaml:"port,omitempty"`
			} `yaml:"group,omitempty"`
		} `yaml:"api,omitempty"`
		Kafka struct {
			Client_id *string `yaml:"client_id,omitempty"`
			Host *string `yaml:"host,omitempty"`
			Port *int64 `yaml:"port,omitempty"`
			Topics struct {
				Topology *string `yaml:"topology,omitempty"`
			} `yaml:"topics,omitempty"`
			Version *string `yaml:"version,omitempty"`
		} `yaml:"kafka,omitempty"`
	} `yaml:"collector,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	Sevone struct {
		Data_volume struct {
			EmptyDir GenericMap `yaml:"emptyDir,omitempty"`
		} `yaml:"data_volume,omitempty"`
		Service_account_credentials_secret *string `yaml:"service_account_credentials_secret,omitempty"`
		Tls_cert_secret_key *string `yaml:"tls_cert_secret_key,omitempty"`
		Tls_cert_secret_name *string `yaml:"tls_cert_secret_name,omitempty"`
		Transport struct {
			Settings struct {
				DE_HOST *string `yaml:"DE_HOST,omitempty"`
				DE_PROJECT_ID *string `yaml:"DE_PROJECT_ID,omitempty"`
				GRPC_PORT *int64 `yaml:"GRPC_PORT,omitempty"`
				PERSIST *string `yaml:"PERSIST,omitempty"`
				TRANSPORT *string `yaml:"TRANSPORT,omitempty"`
			} `yaml:"settings,omitempty"`
			Sevone_auth_secret *string `yaml:"sevone_auth_secret,omitempty"`
		} `yaml:"transport,omitempty"`
	} `yaml:"sevone,omitempty"`
}

type Datadog_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Db_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	LivenessFailureThreshold *int64 `yaml:"livenessFailureThreshold,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		Size *string `yaml:"size,omitempty"`
	} `yaml:"persistence,omitempty"`
	ReplicaCount *string `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Dynatrace_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Elasticsearch_t struct {
	AppVersion *string `yaml:"appVersion,omitempty"`
	Client struct {
		AdditionalJavaOpts *string `yaml:"additionalJavaOpts,omitempty"`
		AntiAffinity *string `yaml:"antiAffinity,omitempty"`
		HeapSize *string `yaml:"heapSize,omitempty"`
		Ingress struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Hosts GenericMap `yaml:"hosts,omitempty"`
			Password *string `yaml:"password,omitempty"`
			Path *string `yaml:"path,omitempty"`
			Tls GenericMap `yaml:"tls,omitempty"`
			User *string `yaml:"user,omitempty"`
		} `yaml:"ingress,omitempty"`
		InitResources GenericMap `yaml:"initResources,omitempty"`
		LoadBalancerIP GenericMap `yaml:"loadBalancerIP,omitempty"`
		LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeAffinity GenericMap `yaml:"nodeAffinity,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodDisruptionBudget struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			MaxUnavailable *string `yaml:"maxUnavailable,omitempty"`
			MinAvailable *int64 `yaml:"minAvailable,omitempty"`
		} `yaml:"podDisruptionBudget,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		Replicas *int64 `yaml:"replicas,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		ServiceAnnotations GenericMap `yaml:"serviceAnnotations,omitempty"`
		ServiceType *string `yaml:"serviceType,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"client,omitempty"`
	Cluster struct {
		AdditionalJavaOpts *string `yaml:"additionalJavaOpts,omitempty"`
		BootstrapShellCommand *string `yaml:"bootstrapShellCommand,omitempty"`
		Config GenericMap `yaml:"config,omitempty"`
		Env GenericMap `yaml:"env,omitempty"`
		KeystoreSecret *string `yaml:"keystoreSecret,omitempty"`
		Name *string `yaml:"name,omitempty"`
		Plugins GenericMap `yaml:"plugins,omitempty"`
		XpackEnable *bool `yaml:"xpackEnable,omitempty"`
	} `yaml:"cluster,omitempty"`
	Data struct {
		AdditionalJavaOpts *string `yaml:"additionalJavaOpts,omitempty"`
		AntiAffinity *string `yaml:"antiAffinity,omitempty"`
		ExposeHttp *bool `yaml:"exposeHttp,omitempty"`
		HeapSize *string `yaml:"heapSize,omitempty"`
		Hooks struct {
			Drain struct {
				Enabled *bool `yaml:"enabled,omitempty"`
			} `yaml:"drain,omitempty"`
		} `yaml:"hooks,omitempty"`
		InitResources GenericMap `yaml:"initResources,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeAffinity GenericMap `yaml:"nodeAffinity,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		Persistence struct {
			AccessMode *string `yaml:"accessMode,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Name *string `yaml:"name,omitempty"`
			Size *string `yaml:"size,omitempty"`
			StorageClass *string `yaml:"storageClass,omitempty"`
		} `yaml:"persistence,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodDisruptionBudget struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			MaxUnavailable *int64 `yaml:"maxUnavailable,omitempty"`
			MinAvailable *string `yaml:"minAvailable,omitempty"`
		} `yaml:"podDisruptionBudget,omitempty"`
		PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReadinessProbe struct {
			HttpGet struct {
				Path *string `yaml:"path,omitempty"`
				Port *int64 `yaml:"port,omitempty"`
			} `yaml:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		} `yaml:"readinessProbe,omitempty"`
		Replicas *int64 `yaml:"replicas,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		TerminationGracePeriodSeconds *int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
		UpdateStrategy struct {
			Type *string `yaml:"type,omitempty"`
		} `yaml:"updateStrategy,omitempty"`
	} `yaml:"data,omitempty"`
	ExtraInitContainers *string `yaml:"extraInitContainers,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	InitImage struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"initImage,omitempty"`
	Master struct {
		AdditionalJavaOpts *string `yaml:"additionalJavaOpts,omitempty"`
		AntiAffinity *string `yaml:"antiAffinity,omitempty"`
		ExposeHttp *bool `yaml:"exposeHttp,omitempty"`
		HeapSize *string `yaml:"heapSize,omitempty"`
		InitResources GenericMap `yaml:"initResources,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeAffinity GenericMap `yaml:"nodeAffinity,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		Persistence struct {
			AccessMode *string `yaml:"accessMode,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Name *string `yaml:"name,omitempty"`
			Size *string `yaml:"size,omitempty"`
			StorageClass *string `yaml:"storageClass,omitempty"`
		} `yaml:"persistence,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodDisruptionBudget struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			MaxUnavailable *string `yaml:"maxUnavailable,omitempty"`
			MinAvailable *int64 `yaml:"minAvailable,omitempty"`
		} `yaml:"podDisruptionBudget,omitempty"`
		PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReadinessProbe struct {
			HttpGet struct {
				Path *string `yaml:"path,omitempty"`
				Port *int64 `yaml:"port,omitempty"`
			} `yaml:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		} `yaml:"readinessProbe,omitempty"`
		Replicas *int64 `yaml:"replicas,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
		UpdateStrategy struct {
			Type *string `yaml:"type,omitempty"`
		} `yaml:"updateStrategy,omitempty"`
	} `yaml:"master,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	PodSecurityPolicy struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"podSecurityPolicy,omitempty"`
	ServiceAccounts struct {
		Client struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"client,omitempty"`
		Data struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"data,omitempty"`
		Master struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"master,omitempty"`
	} `yaml:"serviceAccounts,omitempty"`
	SysctlInitContainer struct {
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"sysctlInitContainer,omitempty"`
	TestFramework struct {
		Image *string `yaml:"image,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"testFramework,omitempty"`
}

type Extractor_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Fluent_bit_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Config struct {
		AutoKubernetesLabels *bool `yaml:"autoKubernetesLabels,omitempty"`
		BatchSize *int64 `yaml:"batchSize,omitempty"`
		BatchWait *int64 `yaml:"batchWait,omitempty"`
		ExtraOutputs GenericMap `yaml:"extraOutputs,omitempty"`
		K8sLoggingParser *string `yaml:"k8sLoggingParser,omitempty"`
		LabelMap struct {
			Kubernetes struct {
				Container_name *string `yaml:"container_name,omitempty"`
				Host *string `yaml:"host,omitempty"`
				Labels struct {
					App *string `yaml:"app,omitempty"`
					Release *string `yaml:"release,omitempty"`
				} `yaml:"labels,omitempty"`
				Namespace_name *string `yaml:"namespace_name,omitempty"`
				Pod_name *string `yaml:"pod_name,omitempty"`
			} `yaml:"kubernetes,omitempty"`
			Stream *string `yaml:"stream,omitempty"`
		} `yaml:"labelMap,omitempty"`
		Labels *string `yaml:"labels,omitempty"`
		LineFormat *string `yaml:"lineFormat,omitempty"`
		Loglevel *string `yaml:"loglevel,omitempty"`
		Parsers GenericMap `yaml:"parsers,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		RemoveKeys *string `yaml:"removeKeys,omitempty"`
		TenantID *string `yaml:"tenantID,omitempty"`
	} `yaml:"config,omitempty"`
	DeploymentStrategy *string `yaml:"deploymentStrategy,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Loki struct {
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		NameOverride *string `yaml:"nameOverride,omitempty"`
		Password *string `yaml:"password,omitempty"`
		ServiceName *string `yaml:"serviceName,omitempty"`
		ServicePath *string `yaml:"servicePath,omitempty"`
		ServicePort *int64 `yaml:"servicePort,omitempty"`
		ServiceScheme *string `yaml:"serviceScheme,omitempty"`
		User *string `yaml:"user,omitempty"`
	} `yaml:"loki,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
		PspEnabled *bool `yaml:"pspEnabled,omitempty"`
	} `yaml:"rbac,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccount struct {
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	ServiceMonitor struct {
		AdditionalLabels GenericMap `yaml:"additionalLabels,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Interval *string `yaml:"interval,omitempty"`
		ScrapeTimeout *string `yaml:"scrapeTimeout,omitempty"`
	} `yaml:"serviceMonitor,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
	VolumeMounts GenericMap `yaml:"volumeMounts,omitempty"`
	Volumes GenericMap `yaml:"volumes,omitempty"`
}

type Gcp_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Global_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Certmanager struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Issuer *string `yaml:"issuer,omitempty"`
	} `yaml:"certmanager,omitempty"`
	CustomImageNames *bool `yaml:"customImageNames,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Elk *bool `yaml:"elk,omitempty"`
	ExternalArangoDBName *string `yaml:"externalArangoDBName,omitempty"`
	ExternalConsul *string `yaml:"externalConsul,omitempty"`
	ExternalDBName *string `yaml:"externalDBName,omitempty"`
	ExternalDbIP *string `yaml:"externalDbIP,omitempty"`
	ExternalDbPort *string `yaml:"externalDbPort,omitempty"`
	ExternalIP *string `yaml:"externalIP,omitempty"`
	ExternalInfluxDBName *string `yaml:"externalInfluxDBName,omitempty"`
	ExternalKafka *string `yaml:"externalKafka,omitempty"`
	ExternalSyslog *string `yaml:"externalSyslog,omitempty"`
	ExternalTimescaleDBIP *string `yaml:"externalTimescaleDBIP,omitempty"`
	ExternalTimescaleDBName *string `yaml:"externalTimescaleDBName,omitempty"`
	ExternalTimescaleDBPort *string `yaml:"externalTimescaleDBPort,omitempty"`
	ImagePassword *string `yaml:"imagePassword,omitempty"`
	ImagePullSecret *string `yaml:"imagePullSecret,omitempty"`
	ImagePullSecrets GenericMap `yaml:"imagePullSecrets,omitempty"`
	ImageRegistry *string `yaml:"imageRegistry,omitempty"`
	ImageUsername *string `yaml:"imageUsername,omitempty"`
	Ingress struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Domain *string `yaml:"domain,omitempty"`
		Secrets []string `yaml:"secrets,omitempty"`
	} `yaml:"ingress,omitempty"`
	JavaBaseOptions *string `yaml:"javaBaseOptions,omitempty"`
	JavaDebugOptions *string `yaml:"javaDebugOptions,omitempty"`
	JavaEnvironmentOptions *string `yaml:"javaEnvironmentOptions,omitempty"`
	JavaMaxRAMPercentage *string `yaml:"javaMaxRAMPercentage,omitempty"`
	JvmType *string `yaml:"jvmType,omitempty"`
	LivenessFailureThreshold *string `yaml:"livenessFailureThreshold,omitempty"`
	LivenessInitialDelaySecs *string `yaml:"livenessInitialDelaySecs,omitempty"`
	LivenessPeriodSecs *string `yaml:"livenessPeriodSecs,omitempty"`
	LivenessSuccessThreshold *string `yaml:"livenessSuccessThreshold,omitempty"`
	LivenessTimeoutSecs *string `yaml:"livenessTimeoutSecs,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	Postgresql struct {
		ExistingSecret *string `yaml:"existingSecret,omitempty"`
		PostgresqlDatabase *string `yaml:"postgresqlDatabase,omitempty"`
		PostgresqlPassword *string `yaml:"postgresqlPassword,omitempty"`
		PostgresqlUsername *string `yaml:"postgresqlUsername,omitempty"`
		ReplicationPassword *string `yaml:"replicationPassword,omitempty"`
		ReplicationUser *string `yaml:"replicationUser,omitempty"`
		ServicePort *string `yaml:"servicePort,omitempty"`
	} `yaml:"postgresql,omitempty"`
	PullPolicy *string `yaml:"pullPolicy,omitempty"`
	ReadinessFailureThreshold *string `yaml:"readinessFailureThreshold,omitempty"`
	ReadinessInitialDelaySecs *string `yaml:"readinessInitialDelaySecs,omitempty"`
	ReadinessPeriodSecs *int64 `yaml:"readinessPeriodSecs,omitempty"`
	ReadinessSuccessThreshold *string `yaml:"readinessSuccessThreshold,omitempty"`
	ReadinessTimeoutSecs *int64 `yaml:"readinessTimeoutSecs,omitempty"`
	Registry *string `yaml:"registry,omitempty"`
	Repository *string `yaml:"repository,omitempty"`
	SecurityContext struct {
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
	} `yaml:"securityContext,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
	StorageClassName *string `yaml:"storageClassName,omitempty"`
	StorageSelector *bool `yaml:"storageSelector,omitempty"`
	Tag *string `yaml:"tag,omitempty"`
	Tolerations []string `yaml:"tolerations,omitempty"`
}

type Grafana_t struct {
	Admin struct {
		DbPasswordKey *string `yaml:"dbPasswordKey,omitempty"`
		ExistingSecret *string `yaml:"existingSecret,omitempty"`
		PasswordKey *string `yaml:"passwordKey,omitempty"`
		UserKey *string `yaml:"userKey,omitempty"`
	} `yaml:"admin,omitempty"`
	AdminPassword *string `yaml:"adminPassword,omitempty"`
	AdminUser *string `yaml:"adminUser,omitempty"`
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Command GenericMap `yaml:"command,omitempty"`
	DashboardProviders struct {
		Dashboardproviders_yaml struct {
			ApiVersion *int64 `yaml:"apiVersion,omitempty"`
		} `yaml:"dashboardproviders.yaml,omitempty"`
	} `yaml:"dashboardProviders,omitempty"`
	Dashboards GenericMap `yaml:"dashboards,omitempty"`
	DashboardsConfigMaps GenericMap `yaml:"dashboardsConfigMaps,omitempty"`
	Datasources struct {
		Datasources_yaml struct {
			ApiVersion *int64 `yaml:"apiVersion,omitempty"`
		} `yaml:"datasources.yaml,omitempty"`
	} `yaml:"datasources,omitempty"`
	DeploymentStrategy struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"deploymentStrategy,omitempty"`
	DownloadDashboards struct {
		Env GenericMap `yaml:"env,omitempty"`
	} `yaml:"downloadDashboards,omitempty"`
	DownloadDashboardsImage struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"downloadDashboardsImage,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	EnvFromSecret *string `yaml:"envFromSecret,omitempty"`
	EnvRenderSecret GenericMap `yaml:"envRenderSecret,omitempty"`
	ExtraConfigmapMounts GenericMap `yaml:"extraConfigmapMounts,omitempty"`
	ExtraContainers *string `yaml:"extraContainers,omitempty"`
	ExtraEmptyDirMounts GenericMap `yaml:"extraEmptyDirMounts,omitempty"`
	ExtraInitContainers GenericMap `yaml:"extraInitContainers,omitempty"`
	ExtraSecretMounts GenericMap `yaml:"extraSecretMounts,omitempty"`
	ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Grafana_ini GenericMap `yaml:"grafana.ini,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Ingress struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExtraPaths *string `yaml:"extraPaths,omitempty"`
		Hosts GenericMap `yaml:"hosts,omitempty"`
		Labels GenericMap `yaml:"labels,omitempty"`
		Path *string `yaml:"path,omitempty"`
		Tls GenericMap `yaml:"tls,omitempty"`
	} `yaml:"ingress,omitempty"`
	InitChownData struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
	} `yaml:"initChownData,omitempty"`
	Labels GenericMap `yaml:"labels,omitempty"`
	Ldap struct {
		Config *string `yaml:"config,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExistingSecret *string `yaml:"existingSecret,omitempty"`
	} `yaml:"ldap,omitempty"`
	LivenessProbe struct {
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *int64 `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	Notifiers GenericMap `yaml:"notifiers,omitempty"`
	Persistence struct {
		AccessModes GenericMap `yaml:"accessModes,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExistingClaim *string `yaml:"existingClaim,omitempty"`
		Finalizers GenericMap `yaml:"finalizers,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClassName *string `yaml:"storageClassName,omitempty"`
		SubPath *string `yaml:"subPath,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"persistence,omitempty"`
	Plugins *string `yaml:"plugins,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodDisruptionBudget struct {
		MaxUnavailable *string `yaml:"maxUnavailable,omitempty"`
		MinAvailable *string `yaml:"minAvailable,omitempty"`
	} `yaml:"podDisruptionBudget,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PodPortName *string `yaml:"podPortName,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
		ExtraClusterRoleRules GenericMap `yaml:"extraClusterRoleRules,omitempty"`
		ExtraRoleRules GenericMap `yaml:"extraRoleRules,omitempty"`
		Namespaced *bool `yaml:"namespaced,omitempty"`
		PspEnabled *bool `yaml:"pspEnabled,omitempty"`
		PspUseAppArmor *bool `yaml:"pspUseAppArmor,omitempty"`
	} `yaml:"rbac,omitempty"`
	ReadinessProbe struct {
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *int64 `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	Replicas *int64 `yaml:"replicas,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	SchedulerName *string `yaml:"schedulerName,omitempty"`
	SecurityContext struct {
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Service struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		ClusterIP *string `yaml:"clusterIP,omitempty"`
		Labels GenericMap `yaml:"labels,omitempty"`
		LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
		LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
		NodePort *string `yaml:"nodePort,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		PortName *string `yaml:"portName,omitempty"`
		TargetPort *int64 `yaml:"targetPort,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceAccount struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NameTest *string `yaml:"nameTest,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	Sidecar struct {
		Dashboards struct {
			SCProvider *bool `yaml:"SCProvider,omitempty"`
			DefaultFolderName GenericMap `yaml:"defaultFolderName,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Folder *string `yaml:"folder,omitempty"`
			Label *string `yaml:"label,omitempty"`
			Provider struct {
				DisableDelete *bool `yaml:"disableDelete,omitempty"`
				Folder *string `yaml:"folder,omitempty"`
				Name *string `yaml:"name,omitempty"`
				Orgid *int64 `yaml:"orgid,omitempty"`
				Type *string `yaml:"type,omitempty"`
			} `yaml:"provider,omitempty"`
			SearchNamespace *string `yaml:"searchNamespace,omitempty"`
		} `yaml:"dashboards,omitempty"`
		Datasources struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			Label *string `yaml:"label,omitempty"`
			SearchNamespace *string `yaml:"searchNamespace,omitempty"`
		} `yaml:"datasources,omitempty"`
		Image *string `yaml:"image,omitempty"`
		ImagePullPolicy *string `yaml:"imagePullPolicy,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SkipTlsVerify *string `yaml:"skipTlsVerify,omitempty"`
	} `yaml:"sidecar,omitempty"`
	Smtp struct {
		ExistingSecret *string `yaml:"existingSecret,omitempty"`
		PasswordKey *string `yaml:"passwordKey,omitempty"`
		UserKey *string `yaml:"userKey,omitempty"`
	} `yaml:"smtp,omitempty"`
	TestFramework struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Image *string `yaml:"image,omitempty"`
		SecurityContext GenericMap `yaml:"securityContext,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"testFramework,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
}

type Group_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Hds_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type History_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	LivenessFailureThreshold *int64 `yaml:"livenessFailureThreshold,omitempty"`
	LivenessTimeoutSecs *int64 `yaml:"livenessTimeoutSecs,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReadinessTimeoutSecs *int64 `yaml:"readinessTimeoutSecs,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Horizon_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Hpe3par_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Hyperflex_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Hyperv_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Influxdb_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		Size *string `yaml:"size,omitempty"`
	} `yaml:"persistence,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Intersight_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Intersight_integration_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Istio_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Istioingress_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
	Routes struct {
		Match []string `yaml:"match,omitempty"`
		Route []string `yaml:"route,omitempty"`
	} `yaml:"routes,omitempty"`
}

type Jaeger_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Kafka_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Java struct {
		Options *string `yaml:"options,omitempty"`
	} `yaml:"java,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		Size *string `yaml:"size,omitempty"`
	} `yaml:"persistence,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Kibana_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	AuthProxyEnabled *bool `yaml:"authProxyEnabled,omitempty"`
	Commandline struct {
		Args GenericMap `yaml:"args,omitempty"`
	} `yaml:"commandline,omitempty"`
	DashboardImport struct {
		Dashboards GenericMap `yaml:"dashboards,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Timeout *int64 `yaml:"timeout,omitempty"`
		Xpackauth struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			Password *string `yaml:"password,omitempty"`
			Username *string `yaml:"username,omitempty"`
		} `yaml:"xpackauth,omitempty"`
	} `yaml:"dashboardImport,omitempty"`
	Deployment struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
	} `yaml:"deployment,omitempty"`
	Elasticsearch struct {
		Host *string `yaml:"host,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
	} `yaml:"elasticsearch,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	ExtraConfigMapMounts GenericMap `yaml:"extraConfigMapMounts,omitempty"`
	ExtraContainers *string `yaml:"extraContainers,omitempty"`
	ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
	ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
	Files *string `yaml:"files,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Ingress struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Hosts GenericMap `yaml:"hosts,omitempty"`
		Tls GenericMap `yaml:"tls,omitempty"`
	} `yaml:"ingress,omitempty"`
	InitContainers GenericMap `yaml:"initContainers,omitempty"`
	LivenessProbe struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		Path *string `yaml:"path,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	PersistentVolumeClaim struct {
		AccessModes GenericMap `yaml:"accessModes,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExistingClaim *bool `yaml:"existingClaim,omitempty"`
		Name *string `yaml:"name,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClass *string `yaml:"storageClass,omitempty"`
	} `yaml:"persistentVolumeClaim,omitempty"`
	Plugins struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Reset *bool `yaml:"reset,omitempty"`
		Values GenericMap `yaml:"values,omitempty"`
	} `yaml:"plugins,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	ReadinessProbe struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		Path *string `yaml:"path,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	RevisionHistoryLimit *int64 `yaml:"revisionHistoryLimit,omitempty"`
	SecurityContext struct {
		AllowPrivilegeEscalation *bool `yaml:"allowPrivilegeEscalation,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Service struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		AuthProxyPort *string `yaml:"authProxyPort,omitempty"`
		ClusterIP *string `yaml:"clusterIP,omitempty"`
		ExternalPort *int64 `yaml:"externalPort,omitempty"`
		InternalPort *int64 `yaml:"internalPort,omitempty"`
		Labels GenericMap `yaml:"labels,omitempty"`
		LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
		LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
		NodePort *string `yaml:"nodePort,omitempty"`
		PortName *string `yaml:"portName,omitempty"`
		Selector GenericMap `yaml:"selector,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceAccount struct {
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	TestFramework struct {
		Image *string `yaml:"image,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"testFramework,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
}

type Kubeturbo_t struct {
	HANodeConfig struct {
		NodeRoles *string `yaml:"nodeRoles,omitempty"`
	} `yaml:"HANodeConfig,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Args struct {
		Kubelethttps *bool `yaml:"kubelethttps,omitempty"`
		Kubeletport *int64 `yaml:"kubeletport,omitempty"`
		Logginglevel *int64 `yaml:"logginglevel,omitempty"`
		Pre16k8sVersion *bool `yaml:"pre16k8sVersion,omitempty"`
		Stitchuuid *bool `yaml:"stitchuuid,omitempty"`
	} `yaml:"args,omitempty"`
	DaemonPodDetectors *string `yaml:"daemonPodDetectors,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	RestAPIConfig struct {
		OpsManagerPassword *string `yaml:"opsManagerPassword,omitempty"`
		OpsManagerUserName *string `yaml:"opsManagerUserName,omitempty"`
		TurbonomicCredentialsSecretName *string `yaml:"turbonomicCredentialsSecretName,omitempty"`
	} `yaml:"restAPIConfig,omitempty"`
	RoleName *string `yaml:"roleName,omitempty"`
	ServerMeta struct {
		Proxy *string `yaml:"proxy,omitempty"`
		TurboServer *string `yaml:"turboServer,omitempty"`
		Version *string `yaml:"version,omitempty"`
	} `yaml:"serverMeta,omitempty"`
	TargetConfig struct {
		TargetName *string `yaml:"targetName,omitempty"`
		TargetType *string `yaml:"targetType,omitempty"`
	} `yaml:"targetConfig,omitempty"`
}

type Logstash_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Config struct {
		Config_reload_automatic *string `yaml:"config.reload.automatic,omitempty"`
		Path_config *string `yaml:"path.config,omitempty"`
		Path_data *string `yaml:"path.data,omitempty"`
		Queue_checkpoint_writes *int64 `yaml:"queue.checkpoint.writes,omitempty"`
		Queue_drain *string `yaml:"queue.drain,omitempty"`
		Queue_max_bytes *string `yaml:"queue.max_bytes,omitempty"`
		Queue_type *string `yaml:"queue.type,omitempty"`
	} `yaml:"config,omitempty"`
	Elasticsearch struct {
		Host *string `yaml:"host,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
	} `yaml:"elasticsearch,omitempty"`
	Exporter struct {
		Logstash struct {
			Config GenericMap `yaml:"config,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Env GenericMap `yaml:"env,omitempty"`
			Image struct {
				PullPolicy *string `yaml:"pullPolicy,omitempty"`
				Repository *string `yaml:"repository,omitempty"`
				Tag *string `yaml:"tag,omitempty"`
			} `yaml:"image,omitempty"`
			LivenessProbe struct {
				FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
				HttpGet struct {
					Path *string `yaml:"path,omitempty"`
					Port *string `yaml:"port,omitempty"`
				} `yaml:"httpGet,omitempty"`
				PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
				SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
				TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
			} `yaml:"livenessProbe,omitempty"`
			Path *string `yaml:"path,omitempty"`
			Port *int64 `yaml:"port,omitempty"`
			ReadinessProbe struct {
				FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
				HttpGet struct {
					Path *string `yaml:"path,omitempty"`
					Port *string `yaml:"port,omitempty"`
				} `yaml:"httpGet,omitempty"`
				PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
				SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
				TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
			} `yaml:"readinessProbe,omitempty"`
			Resources struct {
				Limits struct {
					Cpu *string `yaml:"cpu,omitempty"`
					Memory *string `yaml:"memory,omitempty"`
				} `yaml:"limits,omitempty"`
				Requests struct {
					Cpu *string `yaml:"cpu,omitempty"`
					Memory *string `yaml:"memory,omitempty"`
				} `yaml:"requests,omitempty"`
			} `yaml:"resources,omitempty"`
			Target struct {
				Path *string `yaml:"path,omitempty"`
				Port *int64 `yaml:"port,omitempty"`
			} `yaml:"target,omitempty"`
		} `yaml:"logstash,omitempty"`
	} `yaml:"exporter,omitempty"`
	Files GenericMap `yaml:"files,omitempty"`
	Filters struct {
		Main *string `yaml:"main,omitempty"`
	} `yaml:"filters,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Ingress struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Hosts GenericMap `yaml:"hosts,omitempty"`
		Path *string `yaml:"path,omitempty"`
		Tls GenericMap `yaml:"tls,omitempty"`
	} `yaml:"ingress,omitempty"`
	Inputs struct {
		Main *string `yaml:"main,omitempty"`
	} `yaml:"inputs,omitempty"`
	LivenessProbe struct {
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *string `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	Outputs struct {
		Main *string `yaml:"main,omitempty"`
	} `yaml:"outputs,omitempty"`
	Patterns GenericMap `yaml:"patterns,omitempty"`
	Persistence struct {
		AccessMode *string `yaml:"accessMode,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClass *string `yaml:"storageClass,omitempty"`
	} `yaml:"persistence,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodDisruptionBudget struct {
		MaxUnavailable *int64 `yaml:"maxUnavailable,omitempty"`
	} `yaml:"podDisruptionBudget,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	Ports GenericMap `yaml:"ports,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	ReadinessProbe struct {
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *string `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	SecurityContext struct {
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Service struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		ClusterIP *string `yaml:"clusterIP,omitempty"`
		ExternalTrafficPolicy *string `yaml:"externalTrafficPolicy,omitempty"`
		LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
		NodePort *string `yaml:"nodePort,omitempty"`
		Ports struct {
			Syslog_tcp struct {
				Port *int64 `yaml:"port,omitempty"`
				Protocol *string `yaml:"protocol,omitempty"`
				TargetPort *string `yaml:"targetPort,omitempty"`
			} `yaml:"syslog-tcp,omitempty"`
		} `yaml:"ports,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	TerminationGracePeriodSeconds *int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
	UpdateStrategy struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"updateStrategy,omitempty"`
	VolumeMounts GenericMap `yaml:"volumeMounts,omitempty"`
	Volumes GenericMap `yaml:"volumes,omitempty"`
}

type Loki_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Client struct {
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NameOverride *string `yaml:"nameOverride,omitempty"`
	} `yaml:"client,omitempty"`
	Config struct {
		Auth_enabled *bool `yaml:"auth_enabled,omitempty"`
		Chunk_store_config struct {
			Max_look_back_period *string `yaml:"max_look_back_period,omitempty"`
		} `yaml:"chunk_store_config,omitempty"`
		Ingester struct {
			Chunk_block_size *int64 `yaml:"chunk_block_size,omitempty"`
			Chunk_idle_period *string `yaml:"chunk_idle_period,omitempty"`
			Chunk_retain_period *string `yaml:"chunk_retain_period,omitempty"`
			Lifecycler struct {
				Ring struct {
					Kvstore struct {
						Store *string `yaml:"store,omitempty"`
					} `yaml:"kvstore,omitempty"`
					Replication_factor *int64 `yaml:"replication_factor,omitempty"`
				} `yaml:"ring,omitempty"`
			} `yaml:"lifecycler,omitempty"`
			Max_transfer_retries *int64 `yaml:"max_transfer_retries,omitempty"`
		} `yaml:"ingester,omitempty"`
		Limits_config struct {
			Enforce_metric_name *bool `yaml:"enforce_metric_name,omitempty"`
			Reject_old_samples *bool `yaml:"reject_old_samples,omitempty"`
			Reject_old_samples_max_age *string `yaml:"reject_old_samples_max_age,omitempty"`
		} `yaml:"limits_config,omitempty"`
		Schema_config GenericMap `yaml:"schema_config,omitempty"`
		Server struct {
			Http_listen_port *int64 `yaml:"http_listen_port,omitempty"`
		} `yaml:"server,omitempty"`
		Storage_config struct {
			Boltdb struct {
				Directory *string `yaml:"directory,omitempty"`
			} `yaml:"boltdb,omitempty"`
			Filesystem struct {
				Directory *string `yaml:"directory,omitempty"`
			} `yaml:"filesystem,omitempty"`
		} `yaml:"storage_config,omitempty"`
		Table_manager struct {
			Retention_deletes_enabled *bool `yaml:"retention_deletes_enabled,omitempty"`
			Retention_period *string `yaml:"retention_period,omitempty"`
		} `yaml:"table_manager,omitempty"`
	} `yaml:"config,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
	ExtraContainers GenericMap `yaml:"extraContainers,omitempty"`
	ExtraPorts GenericMap `yaml:"extraPorts,omitempty"`
	ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
	ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	InitContainers GenericMap `yaml:"initContainers,omitempty"`
	LivenessProbe struct {
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *string `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NetworkPolicy struct {
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"networkPolicy,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	Persistence struct {
		AccessModes GenericMap `yaml:"accessModes,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExistingClaim *string `yaml:"existingClaim,omitempty"`
		Selector GenericMap `yaml:"selector,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClassName *string `yaml:"storageClassName,omitempty"`
		SubPath *string `yaml:"subPath,omitempty"`
	} `yaml:"persistence,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodDisruptionBudget GenericMap `yaml:"podDisruptionBudget,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
		PspEnabled *bool `yaml:"pspEnabled,omitempty"`
	} `yaml:"rbac,omitempty"`
	ReadinessProbe struct {
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *string `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	Replicas *int64 `yaml:"replicas,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	SecurityContext struct {
		FsGroup *int64 `yaml:"fsGroup,omitempty"`
		RunAsGroup *int64 `yaml:"runAsGroup,omitempty"`
		RunAsNonRoot *bool `yaml:"runAsNonRoot,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Service struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		ClusterIP *string `yaml:"clusterIP,omitempty"`
		Labels GenericMap `yaml:"labels,omitempty"`
		NodePort *string `yaml:"nodePort,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceAccount struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	ServiceMonitor struct {
		AdditionalLabels GenericMap `yaml:"additionalLabels,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Interval *string `yaml:"interval,omitempty"`
		ScrapeTimeout *string `yaml:"scrapeTimeout,omitempty"`
	} `yaml:"serviceMonitor,omitempty"`
	TerminationGracePeriodSeconds *int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
	Tracing struct {
		JaegerAgentHost *string `yaml:"jaegerAgentHost,omitempty"`
	} `yaml:"tracing,omitempty"`
	UpdateStrategy struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"updateStrategy,omitempty"`
}

type Market_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_acims_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_actionscript_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_aix_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_apic_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_appdynamics_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_appinsights_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_aws_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_awsbilling_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_awscost_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_awslambda_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_azure_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_azurecost_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_azureea_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_azuresp_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_azurevolumes_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_baremetal_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_cloudfoundry_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_compellent_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_customdata_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_datadog_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_dynatrace_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_gcp_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_gcpcost_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_hds_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_horizon_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_hpe3par_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_hyperflex_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_hyperv_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_intersight_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_intersighthyperflex_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_intersightucs_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_istio_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_mssql_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_mysql_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_netapp_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_netflow_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_newrelic_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_nutanix_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_oneview_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_openstack_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_pivotal_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_pure_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_rhv_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_scaleio_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_servicenow_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_snmp_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_terraform_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_tetration_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_tomcat_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_ucs_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_ucsdirector_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_udt_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vcd_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vcenter_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vcenterbrowsing_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vmax_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vmm_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_vplex_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_wmi_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mediation_xtremio_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Metron_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Ml_datastore_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Ml_training_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Mssql_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Mysql_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Netapp_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Netflow_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Newrelic_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Nginx_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Nginxingress_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Nutanix_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Oneview_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Openshiftingress_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Openstack_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Pivotal_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Plan_orchestrator_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Platform_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Postgresql_t struct {
	ConfigurationConfigMap *string `yaml:"configurationConfigMap,omitempty"`
	ExistingSecret *string `yaml:"existingSecret,omitempty"`
	ExtendedConfConfigMap *string `yaml:"extendedConfConfigMap,omitempty"`
	ExtraEnv GenericMap `yaml:"extraEnv,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global struct {
		Postgresql GenericMap `yaml:"postgresql,omitempty"`
	} `yaml:"global,omitempty"`
	Image struct {
		Debug *bool `yaml:"debug,omitempty"`
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Registry *string `yaml:"registry,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	InitdbScripts GenericMap `yaml:"initdbScripts,omitempty"`
	InitdbScriptsConfigMap *string `yaml:"initdbScriptsConfigMap,omitempty"`
	InitdbScriptsSecret *string `yaml:"initdbScriptsSecret,omitempty"`
	LivenessProbe struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"livenessProbe,omitempty"`
	Master struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
		ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodLabels GenericMap `yaml:"podLabels,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"master,omitempty"`
	Metrics struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
			Registry *string `yaml:"registry,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		LivenessProbe struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
			InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
			PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
			SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
			TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
		} `yaml:"livenessProbe,omitempty"`
		ReadinessProbe struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
			InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
			PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
			SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
			TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
		} `yaml:"readinessProbe,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SecurityContext struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		ServiceMonitor struct {
			AdditionalLabels GenericMap `yaml:"additionalLabels,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Interval *string `yaml:"interval,omitempty"`
			Namespace *string `yaml:"namespace,omitempty"`
			ScrapeTimeout *string `yaml:"scrapeTimeout,omitempty"`
		} `yaml:"serviceMonitor,omitempty"`
	} `yaml:"metrics,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NetworkPolicy struct {
		AllowExternal *bool `yaml:"allowExternal,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"networkPolicy,omitempty"`
	Persistence struct {
		AccessModes GenericMap `yaml:"accessModes,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExistingClaim GenericMap `yaml:"existingClaim,omitempty"`
		MountPath *string `yaml:"mountPath,omitempty"`
		Size *string `yaml:"size,omitempty"`
		StorageClass *string `yaml:"storageClass,omitempty"`
		SubPath *string `yaml:"subPath,omitempty"`
	} `yaml:"persistence,omitempty"`
	PgHbaConfiguration *string `yaml:"pgHbaConfiguration,omitempty"`
	PostgresqlConfiguration *string `yaml:"postgresqlConfiguration,omitempty"`
	PostgresqlDataDir *string `yaml:"postgresqlDataDir,omitempty"`
	PostgresqlDatabase *string `yaml:"postgresqlDatabase,omitempty"`
	PostgresqlExtendedConf GenericMap `yaml:"postgresqlExtendedConf,omitempty"`
	PostgresqlInitdbArgs *string `yaml:"postgresqlInitdbArgs,omitempty"`
	PostgresqlInitdbWalDir *string `yaml:"postgresqlInitdbWalDir,omitempty"`
	PostgresqlPassword *string `yaml:"postgresqlPassword,omitempty"`
	PostgresqlUsername *string `yaml:"postgresqlUsername,omitempty"`
	ReadinessProbe struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	Replication struct {
		ApplicationName *string `yaml:"applicationName,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		NumSynchronousReplicas *int64 `yaml:"numSynchronousReplicas,omitempty"`
		Password *string `yaml:"password,omitempty"`
		SlaveReplicas *int64 `yaml:"slaveReplicas,omitempty"`
		SynchronousCommit *string `yaml:"synchronousCommit,omitempty"`
		User *string `yaml:"user,omitempty"`
	} `yaml:"replication,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	SchedulerName *string `yaml:"schedulerName,omitempty"`
	SecurityContext struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	Service struct {
		Annotations GenericMap `yaml:"annotations,omitempty"`
		ClusterIP *string `yaml:"clusterIP,omitempty"`
		LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
		LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
		NodePort *string `yaml:"nodePort,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceAccount struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	Slave struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
		ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodLabels GenericMap `yaml:"podLabels,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"slave,omitempty"`
	TerminationGracePeriodSeconds *string `yaml:"terminationGracePeriodSeconds,omitempty"`
	UpdateStrategy struct {
		Type *string `yaml:"type,omitempty"`
	} `yaml:"updateStrategy,omitempty"`
	UsePasswordFile *string `yaml:"usePasswordFile,omitempty"`
	VolumePermissions struct {
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
			Registry *string `yaml:"registry,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		SecurityContext struct {
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
	} `yaml:"volumePermissions,omitempty"`
}

type Prometheus_t struct {
	AlertRelabelConfigs *string `yaml:"alertRelabelConfigs,omitempty"`
	Alertmanager struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		BaseURL *string `yaml:"baseURL,omitempty"`
		ConfigFileName *string `yaml:"configFileName,omitempty"`
		ConfigFromSecret *string `yaml:"configFromSecret,omitempty"`
		ConfigMapOverrideName *string `yaml:"configMapOverrideName,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
		ExtraEnv GenericMap `yaml:"extraEnv,omitempty"`
		ExtraSecretMounts GenericMap `yaml:"extraSecretMounts,omitempty"`
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Ingress struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExtraLabels GenericMap `yaml:"extraLabels,omitempty"`
			ExtraPaths *string `yaml:"extraPaths,omitempty"`
			Hosts GenericMap `yaml:"hosts,omitempty"`
			Tls GenericMap `yaml:"tls,omitempty"`
		} `yaml:"ingress,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PersistentVolume struct {
			AccessModes GenericMap `yaml:"accessModes,omitempty"`
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExistingClaim *string `yaml:"existingClaim,omitempty"`
			MountPath *string `yaml:"mountPath,omitempty"`
			Size *string `yaml:"size,omitempty"`
			StorageClass *string `yaml:"storageClass,omitempty"`
			SubPath *string `yaml:"subPath,omitempty"`
		} `yaml:"persistentVolume,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodSecurityPolicy struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
		} `yaml:"podSecurityPolicy,omitempty"`
		PrefixURL *string `yaml:"prefixURL,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SchedulerName *string `yaml:"schedulerName,omitempty"`
		SecurityContext struct {
			FsGroup *int64 `yaml:"fsGroup,omitempty"`
			RunAsGroup *int64 `yaml:"runAsGroup,omitempty"`
			RunAsNonRoot *bool `yaml:"runAsNonRoot,omitempty"`
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			ClusterIP *string `yaml:"clusterIP,omitempty"`
			EnableMeshPeer *string `yaml:"enableMeshPeer,omitempty"`
			ExternalIPs GenericMap `yaml:"externalIPs,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
			NodePort *string `yaml:"nodePort,omitempty"`
			ServicePort *int64 `yaml:"servicePort,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		StatefulSet struct {
			Enabled *bool `yaml:"enabled,omitempty"`
			Headless struct {
				Annotations GenericMap `yaml:"annotations,omitempty"`
				EnableMeshPeer *string `yaml:"enableMeshPeer,omitempty"`
				Labels GenericMap `yaml:"labels,omitempty"`
				ServicePort *int64 `yaml:"servicePort,omitempty"`
			} `yaml:"headless,omitempty"`
			PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
		} `yaml:"statefulSet,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"alertmanager,omitempty"`
	AlertmanagerFiles struct {
		Alertmanager_yml struct {
			Global GenericMap `yaml:"global,omitempty"`
			Route struct {
				Group_interval *string `yaml:"group_interval,omitempty"`
				Group_wait *string `yaml:"group_wait,omitempty"`
				Receiver *string `yaml:"receiver,omitempty"`
				Repeat_interval *string `yaml:"repeat_interval,omitempty"`
			} `yaml:"route,omitempty"`
		} `yaml:"alertmanager.yml,omitempty"`
	} `yaml:"alertmanagerFiles,omitempty"`
	ConfigmapReload struct {
		ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
		ExtraConfigmapMounts GenericMap `yaml:"extraConfigmapMounts,omitempty"`
		ExtraVolumeDirs GenericMap `yaml:"extraVolumeDirs,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Name *string `yaml:"name,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
	} `yaml:"configmapReload,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	ExtraScrapeConfigs *string `yaml:"extraScrapeConfigs,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	ImagePullSecrets GenericMap `yaml:"imagePullSecrets,omitempty"`
	KubeStateMetrics struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		Args GenericMap `yaml:"args,omitempty"`
		DeploymentAnnotations GenericMap `yaml:"deploymentAnnotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		Pod struct {
			Labels GenericMap `yaml:"labels,omitempty"`
		} `yaml:"pod,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodSecurityPolicy struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
		} `yaml:"podSecurityPolicy,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SecurityContext struct {
			RunAsNonRoot *bool `yaml:"runAsNonRoot,omitempty"`
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			ClusterIP *string `yaml:"clusterIP,omitempty"`
			ExternalIPs GenericMap `yaml:"externalIPs,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
			ServicePort *int64 `yaml:"servicePort,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"kubeStateMetrics,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NetworkPolicy struct {
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"networkPolicy,omitempty"`
	NodeExporter struct {
		DeploymentAnnotations GenericMap `yaml:"deploymentAnnotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
		ExtraConfigmapMounts GenericMap `yaml:"extraConfigmapMounts,omitempty"`
		ExtraHostPathMounts GenericMap `yaml:"extraHostPathMounts,omitempty"`
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		HostNetwork *bool `yaml:"hostNetwork,omitempty"`
		HostPID *bool `yaml:"hostPID,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		Pod struct {
			Labels GenericMap `yaml:"labels,omitempty"`
		} `yaml:"pod,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodSecurityPolicy struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *string `yaml:"enabled,omitempty"`
		} `yaml:"podSecurityPolicy,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SecurityContext GenericMap `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			ClusterIP *string `yaml:"clusterIP,omitempty"`
			ExternalIPs GenericMap `yaml:"externalIPs,omitempty"`
			HostPort *int64 `yaml:"hostPort,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
			ServicePort *int64 `yaml:"servicePort,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
		UpdateStrategy struct {
			Type *string `yaml:"type,omitempty"`
		} `yaml:"updateStrategy,omitempty"`
	} `yaml:"nodeExporter,omitempty"`
	PodSecurityPolicy struct {
		Enabled *bool `yaml:"enabled,omitempty"`
	} `yaml:"podSecurityPolicy,omitempty"`
	Pushgateway struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Ingress struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExtraPaths *string `yaml:"extraPaths,omitempty"`
			Hosts GenericMap `yaml:"hosts,omitempty"`
			Tls GenericMap `yaml:"tls,omitempty"`
		} `yaml:"ingress,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PersistentVolume struct {
			AccessModes GenericMap `yaml:"accessModes,omitempty"`
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExistingClaim *string `yaml:"existingClaim,omitempty"`
			MountPath *string `yaml:"mountPath,omitempty"`
			Size *string `yaml:"size,omitempty"`
			StorageClass *string `yaml:"storageClass,omitempty"`
			SubPath *string `yaml:"subPath,omitempty"`
		} `yaml:"persistentVolume,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodSecurityPolicy struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
		} `yaml:"podSecurityPolicy,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		SecurityContext struct {
			RunAsNonRoot *bool `yaml:"runAsNonRoot,omitempty"`
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			ClusterIP *string `yaml:"clusterIP,omitempty"`
			ExternalIPs GenericMap `yaml:"externalIPs,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
			ServicePort *int64 `yaml:"servicePort,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		Strategy GenericMap `yaml:"strategy,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"pushgateway,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
	} `yaml:"rbac,omitempty"`
	SchedulerName *string `yaml:"schedulerName,omitempty"`
	Server struct {
		Affinity GenericMap `yaml:"affinity,omitempty"`
		BaseURL *string `yaml:"baseURL,omitempty"`
		ConfigMapOverrideName *string `yaml:"configMapOverrideName,omitempty"`
		ConfigPath *string `yaml:"configPath,omitempty"`
		DeploymentAnnotations GenericMap `yaml:"deploymentAnnotations,omitempty"`
		EmptyDir struct {
			SizeLimit *string `yaml:"sizeLimit,omitempty"`
		} `yaml:"emptyDir,omitempty"`
		EnableAdminApi *bool `yaml:"enableAdminApi,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Env GenericMap `yaml:"env,omitempty"`
		ExtraArgs GenericMap `yaml:"extraArgs,omitempty"`
		ExtraConfigmapMounts GenericMap `yaml:"extraConfigmapMounts,omitempty"`
		ExtraHostPathMounts GenericMap `yaml:"extraHostPathMounts,omitempty"`
		ExtraInitContainers GenericMap `yaml:"extraInitContainers,omitempty"`
		ExtraSecretMounts GenericMap `yaml:"extraSecretMounts,omitempty"`
		ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
		ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		Global struct {
			Evaluation_interval *string `yaml:"evaluation_interval,omitempty"`
			Scrape_interval *string `yaml:"scrape_interval,omitempty"`
			Scrape_timeout *string `yaml:"scrape_timeout,omitempty"`
		} `yaml:"global,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repository *string `yaml:"repository,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		Ingress struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExtraLabels GenericMap `yaml:"extraLabels,omitempty"`
			ExtraPaths *string `yaml:"extraPaths,omitempty"`
			Hosts GenericMap `yaml:"hosts,omitempty"`
			Tls GenericMap `yaml:"tls,omitempty"`
		} `yaml:"ingress,omitempty"`
		LivenessProbeInitialDelay *int64 `yaml:"livenessProbeInitialDelay,omitempty"`
		LivenessProbeTimeout *int64 `yaml:"livenessProbeTimeout,omitempty"`
		Name *string `yaml:"name,omitempty"`
		NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
		PersistentVolume struct {
			AccessModes GenericMap `yaml:"accessModes,omitempty"`
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			ExistingClaim *string `yaml:"existingClaim,omitempty"`
			MountPath *string `yaml:"mountPath,omitempty"`
			Size *string `yaml:"size,omitempty"`
			StorageClass *string `yaml:"storageClass,omitempty"`
			SubPath *string `yaml:"subPath,omitempty"`
		} `yaml:"persistentVolume,omitempty"`
		PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
		PodLabels GenericMap `yaml:"podLabels,omitempty"`
		PodSecurityPolicy struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
		} `yaml:"podSecurityPolicy,omitempty"`
		PrefixURL *string `yaml:"prefixURL,omitempty"`
		PriorityClassName *string `yaml:"priorityClassName,omitempty"`
		ReadinessProbeInitialDelay *int64 `yaml:"readinessProbeInitialDelay,omitempty"`
		ReadinessProbeTimeout *int64 `yaml:"readinessProbeTimeout,omitempty"`
		ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
		Resources struct {
			Limits struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"limits,omitempty"`
			Requests struct {
				Cpu *string `yaml:"cpu,omitempty"`
				Memory *string `yaml:"memory,omitempty"`
			} `yaml:"requests,omitempty"`
		} `yaml:"resources,omitempty"`
		Retention *string `yaml:"retention,omitempty"`
		SchedulerName *string `yaml:"schedulerName,omitempty"`
		SecurityContext struct {
			FsGroup *int64 `yaml:"fsGroup,omitempty"`
			RunAsGroup *int64 `yaml:"runAsGroup,omitempty"`
			RunAsNonRoot *bool `yaml:"runAsNonRoot,omitempty"`
			RunAsUser *int64 `yaml:"runAsUser,omitempty"`
		} `yaml:"securityContext,omitempty"`
		Service struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			ClusterIP *string `yaml:"clusterIP,omitempty"`
			ExternalIPs GenericMap `yaml:"externalIPs,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			LoadBalancerIP *string `yaml:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges GenericMap `yaml:"loadBalancerSourceRanges,omitempty"`
			NodePort *string `yaml:"nodePort,omitempty"`
			ServicePort *int64 `yaml:"servicePort,omitempty"`
			Type *string `yaml:"type,omitempty"`
		} `yaml:"service,omitempty"`
		SidecarContainers GenericMap `yaml:"sidecarContainers,omitempty"`
		SkipTSDBLock *bool `yaml:"skipTSDBLock,omitempty"`
		StatefulSet struct {
			Annotations GenericMap `yaml:"annotations,omitempty"`
			Enabled *bool `yaml:"enabled,omitempty"`
			Headless struct {
				Annotations GenericMap `yaml:"annotations,omitempty"`
				Labels GenericMap `yaml:"labels,omitempty"`
				ServicePort *int64 `yaml:"servicePort,omitempty"`
			} `yaml:"headless,omitempty"`
			Labels GenericMap `yaml:"labels,omitempty"`
			PodManagementPolicy *string `yaml:"podManagementPolicy,omitempty"`
		} `yaml:"statefulSet,omitempty"`
		Strategy GenericMap `yaml:"strategy,omitempty"`
		TerminationGracePeriodSeconds *int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
		Tolerations GenericMap `yaml:"tolerations,omitempty"`
	} `yaml:"server,omitempty"`
	ServerFiles struct {
		Alerts GenericMap `yaml:"alerts,omitempty"`
		Prometheus_yml GenericMap `yaml:"prometheus.yml,omitempty"`
		Rules GenericMap `yaml:"rules,omitempty"`
	} `yaml:"serverFiles,omitempty"`
	ServiceAccounts struct {
		Alertmanager struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"alertmanager,omitempty"`
		KubeStateMetrics struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"kubeStateMetrics,omitempty"`
		NodeExporter struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"nodeExporter,omitempty"`
		Pushgateway struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"pushgateway,omitempty"`
		Server struct {
			Create *bool `yaml:"create,omitempty"`
			Name *string `yaml:"name,omitempty"`
		} `yaml:"server,omitempty"`
	} `yaml:"serviceAccounts,omitempty"`
}

type Prometheus_mysql_exporter_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Cloudsqlproxy struct {
		Credentials *string `yaml:"credentials,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Image struct {
			PullPolicy *string `yaml:"pullPolicy,omitempty"`
			Repo *string `yaml:"repo,omitempty"`
			Tag *string `yaml:"tag,omitempty"`
		} `yaml:"image,omitempty"`
		InstanceConnectionName *string `yaml:"instanceConnectionName,omitempty"`
		Port *string `yaml:"port,omitempty"`
	} `yaml:"cloudsqlproxy,omitempty"`
	Collectors GenericMap `yaml:"collectors,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Mysql struct {
		Db *string `yaml:"db,omitempty"`
		ExistingSecret *bool `yaml:"existingSecret,omitempty"`
		Host *string `yaml:"host,omitempty"`
		Param *string `yaml:"param,omitempty"`
		Pass *string `yaml:"pass,omitempty"`
		Port *int64 `yaml:"port,omitempty"`
		Protocol *string `yaml:"protocol,omitempty"`
		User *string `yaml:"user,omitempty"`
	} `yaml:"mysql,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	Service struct {
		ExternalPort *int64 `yaml:"externalPort,omitempty"`
		InternalPort *int64 `yaml:"internalPort,omitempty"`
		Name *string `yaml:"name,omitempty"`
		Type *string `yaml:"type,omitempty"`
	} `yaml:"service,omitempty"`
	ServiceMonitor struct {
		AdditionalLabels GenericMap `yaml:"additionalLabels,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Interval *string `yaml:"interval,omitempty"`
		JobLabel *string `yaml:"jobLabel,omitempty"`
		PodTargetLabels GenericMap `yaml:"podTargetLabels,omitempty"`
		ScrapeTimeout *string `yaml:"scrapeTimeout,omitempty"`
		TargetLabels GenericMap `yaml:"targetLabels,omitempty"`
	} `yaml:"serviceMonitor,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
}

type Prometurbo_t struct {
	Args struct {
		Logginglevel *int64 `yaml:"logginglevel,omitempty"`
	} `yaml:"args,omitempty"`
	BusinessApplications *string `yaml:"businessApplications,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	ExtraPrometheusExporters *string `yaml:"extraPrometheusExporters,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PrometurboTag *string `yaml:"prometurboTag,omitempty"`
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		TurbodifTag *string `yaml:"turbodifTag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	PrometheusServers *string `yaml:"prometheusServers,omitempty"`
	RestAPIConfig struct {
		OpsManagerPassword *string `yaml:"opsManagerPassword,omitempty"`
		OpsManagerUserName *string `yaml:"opsManagerUserName,omitempty"`
		TurbonomicCredentialsSecretName *string `yaml:"turbonomicCredentialsSecretName,omitempty"`
	} `yaml:"restAPIConfig,omitempty"`
	ServerMeta struct {
		Proxy *string `yaml:"proxy,omitempty"`
		TurboServer *string `yaml:"turboServer,omitempty"`
		Version *string `yaml:"version,omitempty"`
	} `yaml:"serverMeta,omitempty"`
	TargetAddress *string `yaml:"targetAddress,omitempty"`
	TargetName *string `yaml:"targetName,omitempty"`
	TargetTypeSuffix *string `yaml:"targetTypeSuffix,omitempty"`
}

type Promtail_t struct {
	Affinity GenericMap `yaml:"affinity,omitempty"`
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Config struct {
		Client struct {
			Backoff_config struct {
				Max_period *string `yaml:"max_period,omitempty"`
				Max_retries *int64 `yaml:"max_retries,omitempty"`
				Min_period *string `yaml:"min_period,omitempty"`
			} `yaml:"backoff_config,omitempty"`
			Batchsize *int64 `yaml:"batchsize,omitempty"`
			Batchwait *string `yaml:"batchwait,omitempty"`
			External_labels GenericMap `yaml:"external_labels,omitempty"`
			Timeout *string `yaml:"timeout,omitempty"`
			Url *string `yaml:"url,omitempty"`
		} `yaml:"client,omitempty"`
		Clients *string `yaml:"clients,omitempty"`
		Positions struct {
			Filename *string `yaml:"filename,omitempty"`
		} `yaml:"positions,omitempty"`
		Server struct {
			Http_listen_port *int64 `yaml:"http_listen_port,omitempty"`
		} `yaml:"server,omitempty"`
		Target_config struct {
			Sync_period *string `yaml:"sync_period,omitempty"`
		} `yaml:"target_config,omitempty"`
	} `yaml:"config,omitempty"`
	DeploymentStrategy *string `yaml:"deploymentStrategy,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	ExtraCommandlineArgs GenericMap `yaml:"extraCommandlineArgs,omitempty"`
	ExtraScrapeConfigs GenericMap `yaml:"extraScrapeConfigs,omitempty"`
	ExtraVolumeMounts GenericMap `yaml:"extraVolumeMounts,omitempty"`
	ExtraVolumes GenericMap `yaml:"extraVolumes,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		PullSecrets GenericMap `yaml:"pullSecrets,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	InitContainer struct {
		Enabled *string `yaml:"enabled,omitempty"`
		FsInotifyMaxUserInstances *string `yaml:"fsInotifyMaxUserInstances,omitempty"`
	} `yaml:"initContainer,omitempty"`
	LivenessProbe GenericMap `yaml:"livenessProbe,omitempty"`
	Loki struct {
		FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
		NameOverride *string `yaml:"nameOverride,omitempty"`
		Password *string `yaml:"password,omitempty"`
		ServiceName *string `yaml:"serviceName,omitempty"`
		ServicePort *int64 `yaml:"servicePort,omitempty"`
		ServiceScheme *string `yaml:"serviceScheme,omitempty"`
		User *string `yaml:"user,omitempty"`
	} `yaml:"loki,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	NodeSelector GenericMap `yaml:"nodeSelector,omitempty"`
	PipelineStages GenericMap `yaml:"pipelineStages,omitempty"`
	PodAnnotations GenericMap `yaml:"podAnnotations,omitempty"`
	PodLabels GenericMap `yaml:"podLabels,omitempty"`
	PriorityClassName *string `yaml:"priorityClassName,omitempty"`
	Rbac struct {
		Create *bool `yaml:"create,omitempty"`
		PspEnabled *bool `yaml:"pspEnabled,omitempty"`
	} `yaml:"rbac,omitempty"`
	ReadinessProbe struct {
		FailureThreshold *int64 `yaml:"failureThreshold,omitempty"`
		HttpGet struct {
			Path *string `yaml:"path,omitempty"`
			Port *string `yaml:"port,omitempty"`
		} `yaml:"httpGet,omitempty"`
		InitialDelaySeconds *int64 `yaml:"initialDelaySeconds,omitempty"`
		PeriodSeconds *int64 `yaml:"periodSeconds,omitempty"`
		SuccessThreshold *int64 `yaml:"successThreshold,omitempty"`
		TimeoutSeconds *int64 `yaml:"timeoutSeconds,omitempty"`
	} `yaml:"readinessProbe,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ScrapeConfigs GenericMap `yaml:"scrapeConfigs,omitempty"`
	SecurityContext struct {
		ReadOnlyRootFilesystem *bool `yaml:"readOnlyRootFilesystem,omitempty"`
		RunAsGroup *int64 `yaml:"runAsGroup,omitempty"`
		RunAsUser *int64 `yaml:"runAsUser,omitempty"`
	} `yaml:"securityContext,omitempty"`
	ServiceAccount struct {
		Create *bool `yaml:"create,omitempty"`
		Name *string `yaml:"name,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	ServiceMonitor struct {
		AdditionalLabels GenericMap `yaml:"additionalLabels,omitempty"`
		Annotations GenericMap `yaml:"annotations,omitempty"`
		Enabled *bool `yaml:"enabled,omitempty"`
		Interval *string `yaml:"interval,omitempty"`
		ScrapeTimeout *string `yaml:"scrapeTimeout,omitempty"`
	} `yaml:"serviceMonitor,omitempty"`
	SyslogService struct {
		Enabled *string `yaml:"enabled,omitempty"`
		Port *string `yaml:"port,omitempty"`
	} `yaml:"syslogService,omitempty"`
	Tolerations GenericMap `yaml:"tolerations,omitempty"`
	VolumeMounts GenericMap `yaml:"volumeMounts,omitempty"`
	Volumes GenericMap `yaml:"volumes,omitempty"`
}

type Properties_t struct {
	Action_orchestrator struct {
		RiskPropagationEnabled *bool `yaml:"riskPropagationEnabled,omitempty"`
	} `yaml:"action-orchestrator,omitempty"`
	Api struct {
		OpenIdAccessTokenUri *string `yaml:"openIdAccessTokenUri,omitempty"`
		OpenIdClientAuthentication *string `yaml:"openIdClientAuthentication,omitempty"`
		OpenIdClientId *string `yaml:"openIdClientId,omitempty"`
		OpenIdClientSecret *string `yaml:"openIdClientSecret,omitempty"`
		OpenIdClients *string `yaml:"openIdClients,omitempty"`
		OpenIdEnabled *bool `yaml:"openIdEnabled,omitempty"`
		OpenIdJwkSetUri *string `yaml:"openIdJwkSetUri,omitempty"`
		OpenIdUserAuthentication *string `yaml:"openIdUserAuthentication,omitempty"`
		OpenIdUserAuthorizationUri *string `yaml:"openIdUserAuthorizationUri,omitempty"`
		OpenIdUserInfoUri *string `yaml:"openIdUserInfoUri,omitempty"`
		SamlEnabled *bool `yaml:"samlEnabled,omitempty"`
		SamlEntityId *string `yaml:"samlEntityId,omitempty"`
		SamlIdpCertificate *string `yaml:"samlIdpCertificate,omitempty"`
		SamlRegistrationId *string `yaml:"samlRegistrationId,omitempty"`
		SamlSpEntityId *string `yaml:"samlSpEntityId,omitempty"`
		SamlWebSsoEndpoint *string `yaml:"samlWebSsoEndpoint,omitempty"`
	} `yaml:"api,omitempty"`
	Auth GenericMap `yaml:"auth,omitempty"`
	Cost GenericMap `yaml:"cost,omitempty"`
	Extractor GenericMap `yaml:"extractor,omitempty"`
	Global struct {
		DbPort *int64 `yaml:"dbPort,omitempty"`
		EnableSecureDBConnection *bool `yaml:"enableSecureDBConnection,omitempty"`
	} `yaml:"global,omitempty"`
	Group GenericMap `yaml:"group,omitempty"`
	History GenericMap `yaml:"history,omitempty"`
	Plan_orchestrator GenericMap `yaml:"plan-orchestrator,omitempty"`
	Repository struct {
		ShowGuestLoad *bool `yaml:"showGuestLoad,omitempty"`
	} `yaml:"repository,omitempty"`
	Topology_processor GenericMap `yaml:"topology-processor,omitempty"`
}

type Pure_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Reporting_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	Enabled *bool `yaml:"enabled,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Repository_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Rhv_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Rsyslog_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Persistence struct {
		Auditsize *string `yaml:"auditsize,omitempty"`
		Logsize *string `yaml:"logsize,omitempty"`
	} `yaml:"persistence,omitempty"`
	ReplicaCount *int64 `yaml:"replicaCount,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Scaleio_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Servicenow_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Snmp_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Terraform_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Tetration_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Timescaledb_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Tomcat_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Topology_processor_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *bool `yaml:"debug,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}

type Training_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Ucs_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Ucsdirector_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Udt_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Vcd_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Vcenter_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Vmax_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Vmm_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Vplex_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Wmi_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Xl_t struct {
	DbSecretName *string `yaml:"dbSecretName,omitempty"`
	Debug *string `yaml:"debug,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Grafana struct {
		AdminPassword *string `yaml:"adminPassword,omitempty"`
	} `yaml:"grafana,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Istioingress struct {
		Enabled *string `yaml:"enabled,omitempty"`
		Routes GenericMap `yaml:"routes,omitempty"`
	} `yaml:"istioingress,omitempty"`
	JavaBaseOptions *string `yaml:"javaBaseOptions,omitempty"`
	JavaComponentOptions *string `yaml:"javaComponentOptions,omitempty"`
	JavaDebugOptions *string `yaml:"javaDebugOptions,omitempty"`
	JavaMaxRAMPercentage *string `yaml:"javaMaxRAMPercentage,omitempty"`
	JavaOptions *string `yaml:"javaOptions,omitempty"`
	JvmType *string `yaml:"jvmType,omitempty"`
	LivenessFailureThreshold *string `yaml:"livenessFailureThreshold,omitempty"`
	LivenessInitialDelaySecs *string `yaml:"livenessInitialDelaySecs,omitempty"`
	LivenessPeriodSecs *string `yaml:"livenessPeriodSecs,omitempty"`
	LivenessSuccessThreshold *string `yaml:"livenessSuccessThreshold,omitempty"`
	LivenessTimeoutSecs *string `yaml:"livenessTimeoutSecs,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Openshiftingress struct {
		Enabled *string `yaml:"enabled,omitempty"`
	} `yaml:"openshiftingress,omitempty"`
	Properties GenericMap `yaml:"properties,omitempty"`
	ReadinessFailureThreshold *string `yaml:"readinessFailureThreshold,omitempty"`
	ReadinessInitialDelaySecs *string `yaml:"readinessInitialDelaySecs,omitempty"`
	ReadinessPeriodSecs *string `yaml:"readinessPeriodSecs,omitempty"`
	ReadinessSuccessThreshold *string `yaml:"readinessSuccessThreshold,omitempty"`
	ReadinessTimeoutSecs *string `yaml:"readinessTimeoutSecs,omitempty"`
	Reporting struct {
		Enabled *string `yaml:"enabled,omitempty"`
	} `yaml:"reporting,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
}

type Xtremio_t struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Zookeeper_t struct {
	Annotations GenericMap `yaml:"annotations,omitempty"`
	Env GenericMap `yaml:"env,omitempty"`
	FullnameOverride *string `yaml:"fullnameOverride,omitempty"`
	Global *string `yaml:"global,omitempty"`
	Image struct {
		PullPolicy *string `yaml:"pullPolicy,omitempty"`
		Repository *string `yaml:"repository,omitempty"`
		ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
		Tag *string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NameOverride *string `yaml:"nameOverride,omitempty"`
	Resources struct {
		Limits struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"limits,omitempty"`
		Requests struct {
			Cpu *string `yaml:"cpu,omitempty"`
			Memory *string `yaml:"memory,omitempty"`
		} `yaml:"requests,omitempty"`
	} `yaml:"resources,omitempty"`
	ServiceAccountName *string `yaml:"serviceAccountName,omitempty"`
	StorageAnnotations GenericMap `yaml:"storageAnnotations,omitempty"`
}
