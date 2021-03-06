package pb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "pb/service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/containers": {
      "post": {
        "summary": "Like 'docker run' command",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"config\":\n    {\n      \"image\": \"nginx\",\n      \"cmd\": [\n        \"-c\", \"printenv \u0026\u0026 ls /usr/share/nginx/html \u0026\u0026 nginx -g \\\"daemon off;\\\"\"\n      ],\n      \"entrypoint\": [\n        \"/bin/bash\"\n      ],\n      \"env\": [\n        \"GOAPTH=/home/vagrant/go\",\n        \"JAVA_HOME=/opt/jdk1.8.0_112\"\n      ],\n      \"exposed_ports\":\n        {\n          \"value\": \n            {\n              \"80\": \"webui\"\n            }\n        },\n      \"volumes\":\n        {\n          \"/etc/nginx/nginx.conf\": \"sysconf\",\n          \"/etc/nginx/conf.d/default.conf\": \"usrconf\",\n          \"/usr/share/nginx/html/\": \"usrdata\"\n        }\n    },\n  \"host_config\":\n    {\n      \"binds\": [\n        \"/home/vagrant/project-php/etc/nginx.conf:/etc/nginx/nginx.conf,Z\",\n        \"/home/vagrant/project-php/etc/default.conf:/etc/nginx/conf.d/default.conf:Z\",\n        \"/home/vagrant/project-php/src/:/usr/share/nginx/html/:ro\"\n      ],\n      \"port_bindings\":\n        {\n          \"value\":\n            {\n              \"80\":\n                {\n                  \"host_port\": \"80\"\n                }\n            }\n        },\n      \"resources\":\n        {\n          \"memory\": 300000000\n        }\n    },\n  \"network_config\":\n    {\n    },\n  \"container_name\": \"nginx\"\n}\nFor output, plus returning elements:\n{\n  \"state_code\": 0, // succeeded, otherwise none zero\n  \"state_message\": \"if failed, provide error information\"\n  \"container_id\": \"returning value of docker engine\"  \n}",
        "operationId": "RunContainer",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerRunData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerRunData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/echo/{value}": {
      "get": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "value",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/network-creation": {
      "post": {
        "operationId": "CreateDockerNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkCreationReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkCreationReqResp"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/networks": {
      "get": {
        "operationId": "ReapDockerNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkData"
            }
          }
        },
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/process-statuses": {
      "post": {
        "summary": "Like 'docker ps' command",
        "operationId": "ProcessStatuses",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerProcessStatusReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerProcessStatusReqResp"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/provisions": {
      "post": {
        "summary": "Run containers with a user namespace information",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"name\": \"fighter and target\"\n  \"namespace\": \"default\"\n  \"metadata\":\n    {\n      \"categroy_name\": \"basic-web-security\",\n      \"class_name\": \"http-protocol\"\n      \"field_name\": \"http-method\"\n    },\n  \"provisionings\": [\n    \u003clist of DockerRunData type, see previous\u003e\n  ]\n}",
        "operationId": "ProvisionContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/pull": {
      "post": {
        "summary": "Like 'docker pull' command",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"image\": \"tomcat:8\"\n}",
        "operationId": "PullImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerPullData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerPullData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/reap-instantiation": {
      "post": {
        "summary": "Find containers with instantiating of a user namespace",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"name\": \"fighter and target\"\n  \"namespace\": \"default\"\n  \"metadata\":\n    {\n      \"categroy_name\": \"default\",\n      \"class_name\": \"default\"\n      \"field_name\": \"default\"\n    }\n}\nFor output, plus a returning object:\n{\n  \"instantiation\": [\n    {\n      // List of Moby Container type\n    }\n  ]\n}",
        "operationId": "ReapInstantiation",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbInstantiationData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbInstantiationData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/reap-registry": {
      "post": {
        "summary": "List registry, include Docker Hub, or private registry (using .docker/config.json)",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"repositories\": [\n    {\n      name: \"127.0.0.1:5000\"\n    }\n  ]\n}\nFor output, plus a returning object:\n{\n  \"repositories\": [\n    {\n      name: \"127.0.0.1:5000\",\n      catalogs: [\n        {\n          \"name\": \"nginx\",\n          \"tags\": [\n            {\n              \"name\": \"latest\"\n            }\n          ]\n        }\n      ]\n    }\n  ],\n  \"state_code\": 0,  // Value greater than zero indicates error \n  \"state_message\": \"...\" // Usually error message \n}",
        "operationId": "ReapRegistryForRepositories",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbRegistryRepositoryData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbRegistryRepositoryData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/sniffing-brns": {
      "post": {
        "operationId": "SniffEtherNetworking",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbEthernetSniffingData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbEthernetSniffingData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/snooping-brns": {
      "post": {
        "operationId": "SnoopBridgedNetworkLandscape",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbBridgedNetworkingData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbBridgedNetworkingData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/terminations": {
      "post": {
        "summary": "Delete containers with a user namespace information",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"name\": \"fighter and target\"\n  \"namespace\": \"default\"\n  \"metadata\":\n    {\n      \"categroy_name\": \"basic-web-security\",\n      \"class_name\": \"http-protocol\"\n      \"field_name\": \"http-method\"\n    }\n}\nAnd returning information append this object for output:\n{\n  \"provisionings\": [\n    list of DockerRunData type, see previous\n  ]\n}",
        "operationId": "TerminationContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "BridgedNetworkingDataContainersNetworkingInfo": {
      "type": "object",
      "properties": {
        "addresses_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BridgedNetworkingDataIPAddressInfo"
          }
        },
        "containers_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyContainerJSON"
          }
        },
        "network_resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyNetworkResource"
          }
        }
      }
    },
    "BridgedNetworkingDataIPAddressInfo": {
      "type": "object",
      "properties": {
        "ipv4": {
          "type": "string",
          "format": "string"
        },
        "ipv6": {
          "type": "string",
          "format": "string"
        },
        "link_info": {
          "$ref": "#/definitions/BridgedNetworkingDataLinkLayerInfo"
        },
        "v4_info": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "v4_lifetime": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "v4_mask": {
          "type": "string",
          "format": "string"
        },
        "v6_info": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "v6_lifetime": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "v6_mask": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "BridgedNetworkingDataLinkLayerInfo": {
      "type": "object",
      "properties": {
        "data_link_conf": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "data_link_ether_brd": {
          "type": "string",
          "format": "string"
        },
        "data_link_ether_mac": {
          "type": "string",
          "format": "string"
        },
        "data_link_frame": {
          "type": "string",
          "format": "string"
        },
        "data_link_netns_id": {
          "type": "string",
          "format": "string"
        },
        "data_link_status": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "index": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "BridgedNetworkingDataLinuxBridgeInfo": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "interfaces": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "ip_addresses_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BridgedNetworkingDataIPAddressInfo"
          }
        },
        "mac_info": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BridgedNetworkingDataLinuxBridgeLearnedMac"
          }
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "stp_enabled": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "BridgedNetworkingDataLinuxBridgeLearnedMac": {
      "type": "object",
      "properties": {
        "ageing_timer": {
          "type": "string",
          "format": "string"
        },
        "is_local": {
          "type": "string",
          "format": "string"
        },
        "mac_addr": {
          "type": "string",
          "format": "string"
        },
        "port_no": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "FiltersArgsValue": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "type": "boolean",
            "format": "boolean"
          }
        }
      }
    },
    "PortMapPortBindings": {
      "type": "object",
      "properties": {
        "port_bindings": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyPortBinding"
          }
        }
      }
    },
    "ProvisioningsDataMetadata": {
      "type": "object",
      "properties": {
        "category_name": {
          "type": "string",
          "format": "string"
        },
        "class_name": {
          "type": "string",
          "format": "string"
        },
        "field_name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "RegistryRepositoryDataCatalog": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RegistryRepositoryDataTag"
          }
        }
      }
    },
    "RegistryRepositoryDataRegistry": {
      "type": "object",
      "properties": {
        "catalogs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RegistryRepositoryDataCatalog"
          }
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "tls_disabled": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "RegistryRepositoryDataTag": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "mobyAddress": {
      "type": "object",
      "properties": {
        "addr": {
          "type": "string",
          "format": "string"
        },
        "prefix_len": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "Address represents an IP address"
    },
    "mobyBindOptions": {
      "type": "object",
      "properties": {
        "propagation": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "BindOptions defines options specific to mounts of type \"bind\".\nto see https://github.com/moby/moby/blob/master/api/types/mount/mount.go"
    },
    "mobyConfig": {
      "type": "object",
      "properties": {
        "args_escaped": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stderr": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stdin": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stdout": {
          "type": "boolean",
          "format": "boolean"
        },
        "cmd": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "domainname": {
          "type": "string",
          "format": "string"
        },
        "entrypoint": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "env": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "exposed_ports": {
          "$ref": "#/definitions/mobyPortSet"
        },
        "healthcheck": {
          "$ref": "#/definitions/mobyHealthConfig"
        },
        "hostname": {
          "type": "string",
          "format": "string"
        },
        "image": {
          "type": "string",
          "format": "string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "network_disabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "on_build": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "open_stdin": {
          "type": "boolean",
          "format": "boolean"
        },
        "shell": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "stdin_once": {
          "type": "boolean",
          "format": "boolean"
        },
        "stop_signal": {
          "type": "string",
          "format": "string"
        },
        "stop_timeout": {
          "type": "integer",
          "format": "int32"
        },
        "tty": {
          "type": "boolean",
          "format": "boolean"
        },
        "user": {
          "type": "string",
          "format": "string"
        },
        "volumes": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "working_dir": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Config contains the configuration data about a container.\nIt should hold only portable information about the container.\nHere, \"portable\" means \"independent from the host we are running on\".\nNon-portable information *should* appear in HostConfig.\nAll fields added to this struct must be marked 'omitempty' to keep getting\npredictable hashes from the old 'v1Compatibility' configuration.\nto see https://github.com/moby/moby/blob/master/api/types/container/config.go"
    },
    "mobyContainer": {
      "type": "object",
      "properties": {
        "Ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyPort"
          }
        },
        "command": {
          "type": "string",
          "format": "string"
        },
        "created": {
          "type": "string",
          "format": "int64"
        },
        "host_config": {
          "$ref": "#/definitions/mobyContainerHostConfig"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "image": {
          "type": "string",
          "format": "string"
        },
        "image_id": {
          "type": "string",
          "format": "string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyMountPoint"
          }
        },
        "names": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "network_settings": {
          "$ref": "#/definitions/mobySummaryNetworkSettings"
        },
        "size_root_fs": {
          "type": "string",
          "format": "int64"
        },
        "size_rw": {
          "type": "string",
          "format": "int64"
        },
        "state": {
          "type": "string",
          "format": "string"
        },
        "status": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Container contains response of Remote API:\nGET  \"/containers/json\""
    },
    "mobyContainerHostConfig": {
      "type": "object",
      "properties": {
        "network_mode": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "mobyContainerJSON": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/mobyConfig"
        },
        "container_json_base": {
          "$ref": "#/definitions/mobyContainerJSONBase"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyMountPoint"
          }
        },
        "network_settings": {
          "$ref": "#/definitions/mobyNetworkSettings"
        }
      },
      "title": "ContainerJSON is newly used struct along with MountPoint"
    },
    "mobyContainerJSONBase": {
      "type": "object",
      "properties": {
        "app_armor_profile": {
          "type": "string",
          "format": "string"
        },
        "args": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "created": {
          "type": "string",
          "format": "string"
        },
        "driver": {
          "type": "string",
          "format": "string"
        },
        "exec_ids": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "graph_driver": {
          "$ref": "#/definitions/mobyGraphDriverData"
        },
        "host_config": {
          "$ref": "#/definitions/mobyHostConfig"
        },
        "hostname_path": {
          "type": "string",
          "format": "string"
        },
        "hosts_path": {
          "type": "string",
          "format": "string"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "image": {
          "type": "string",
          "format": "string"
        },
        "log_path": {
          "type": "string",
          "format": "string"
        },
        "mount_label": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "node": {
          "$ref": "#/definitions/mobyContainerNode"
        },
        "path": {
          "type": "string",
          "format": "string"
        },
        "process_label": {
          "type": "string",
          "format": "string"
        },
        "resolv_conf_path": {
          "type": "string",
          "format": "string"
        },
        "restart_count": {
          "type": "integer",
          "format": "int32"
        },
        "size_root_fs": {
          "type": "string",
          "format": "int64"
        },
        "size_rw": {
          "type": "string",
          "format": "int64"
        },
        "state": {
          "$ref": "#/definitions/mobyContainerState"
        }
      },
      "title": "ContainerJSONBase contains response of Remote API:\nGET \"/containers/{name:.*}/json\""
    },
    "mobyContainerListOptions": {
      "type": "object",
      "properties": {
        "all": {
          "type": "boolean",
          "format": "boolean"
        },
        "before": {
          "type": "string",
          "format": "string"
        },
        "filter": {
          "$ref": "#/definitions/mobyFiltersArgs"
        },
        "latest": {
          "type": "boolean",
          "format": "boolean"
        },
        "limit": {
          "type": "integer",
          "format": "int32"
        },
        "quiet": {
          "type": "boolean",
          "format": "boolean"
        },
        "since": {
          "type": "string",
          "format": "string"
        },
        "size": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "description": "ContainerListOptions holds parameters to list containers with."
    },
    "mobyContainerNode": {
      "type": "object",
      "properties": {
        "addr": {
          "type": "string",
          "format": "string"
        },
        "cpus": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "memory": {
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "ContainerNode stores information about the node that a container\nis running on.  It's only available in Docker Swarm"
    },
    "mobyContainerState": {
      "type": "object",
      "properties": {
        "dead": {
          "type": "boolean",
          "format": "boolean"
        },
        "error": {
          "type": "string",
          "format": "string"
        },
        "exit_code": {
          "type": "integer",
          "format": "int32"
        },
        "finished_at": {
          "type": "string",
          "format": "string"
        },
        "oom_killed": {
          "type": "boolean",
          "format": "boolean"
        },
        "paused": {
          "type": "boolean",
          "format": "boolean"
        },
        "pid": {
          "type": "integer",
          "format": "int32"
        },
        "restarting": {
          "type": "boolean",
          "format": "boolean"
        },
        "running": {
          "type": "boolean",
          "format": "boolean"
        },
        "started_at": {
          "type": "string",
          "format": "string"
        },
        "status": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "ContainerState stores container's running state\nit's part of ContainerJSONBase and will return by \"inspect\" command"
    },
    "mobyDefaultNetworkSettings": {
      "type": "object",
      "properties": {
        "endpoint_id": {
          "type": "string",
          "format": "string"
        },
        "gateway": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ipv6_gateway": {
          "type": "string",
          "format": "string"
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "DefaultNetworkSettings holds network information\nduring the 2 release deprecation period.\nIt will be removed in Docker 1.11."
    },
    "mobyDeviceMapping": {
      "type": "object",
      "properties": {
        "cgroup_permissions": {
          "type": "string",
          "format": "string"
        },
        "path_in_container": {
          "type": "string",
          "format": "string"
        },
        "path_on_host": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "DeviceMapping represents the device mapping between the host and the container."
    },
    "mobyDriverConfig": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "description": "Driver represents a volume driver."
    },
    "mobyEndpointIPAMConfig": {
      "type": "object",
      "properties": {
        "ipv4_address": {
          "type": "string",
          "format": "string"
        },
        "ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "link_local_ips": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "EndpointIPAMConfig represents IPAM configurations for the endpoint"
    },
    "mobyEndpointResource": {
      "type": "object",
      "properties": {
        "endpoint_id": {
          "type": "string",
          "format": "string"
        },
        "ipv4_address": {
          "type": "string",
          "format": "string"
        },
        "ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "EndpointResource contains network resources allocated and used for a container in a network"
    },
    "mobyEndpointSettings": {
      "type": "object",
      "properties": {
        "aliases": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "driver_opts": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "endpoint_id": {
          "type": "string",
          "format": "string"
        },
        "gateway": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ipam_config": {
          "$ref": "#/definitions/mobyEndpointIPAMConfig",
          "title": "Configurations"
        },
        "ipv6_gateway": {
          "type": "string",
          "format": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "network_id": {
          "type": "string",
          "format": "string",
          "title": "Operational data"
        }
      },
      "title": "EndpointSettings stores the network endpoint details"
    },
    "mobyFiltersArgs": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/FiltersArgsValue"
          }
        }
      },
      "title": "Args stores filter arguments as map key:{map key: bool}.\nIt contains an aggregation of the map of arguments (which are in the form\nof -f 'key=value') based on the key, and stores values for the same key\nin a map with string keys and boolean values.\ne.g given -f 'label=label1=1' -f 'label=label2=2' -f 'image.name=ubuntu'\nthe args will be {\"image.name\":{\"ubuntu\":true},\"label\":{\"label1=1\":true,\"label2=2\":true}}"
    },
    "mobyGraphDriverData": {
      "type": "object",
      "properties": {
        "data": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "data\nRequired: true"
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "name\nRequired: true"
        }
      },
      "title": "GraphDriverData Information about a container's graph driver.\nswagger:model GraphDriverData\nto see https://github.com/moby/moby/blob/master/api/types/graph_driver_data.go"
    },
    "mobyHealthConfig": {
      "type": "object",
      "properties": {
        "interval_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Zero means to inherit. Durations are expressed as integer nanoseconds.\nGolang    time.Duration\nInterval is the time to wait between checks."
        },
        "retries": {
          "type": "integer",
          "format": "int32",
          "description": "Retries is the number of consecutive failures needed to consider a container as unhealthy.\nZero means inherit."
        },
        "start_period": {
          "type": "string",
          "format": "int64",
          "description": "Golang time.Duration\nThe start period for the container to initialize before the retries starts to count down."
        },
        "test": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Test is the test to perform to check that the container is healthy.\nAn empty slice means to inherit the default.\nThe options are:\n{} : inherit healthcheck\n{\"NONE\"} : disable healthcheck\n{\"CMD\", args...} : exec arguments directly\n{\"CMD-SHELL\", command} : run command with system's default shell"
        },
        "timeout_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Golang     time.Duration\nTimeout is the time to wait before considering the check to have hung."
        }
      },
      "title": "HealthConfig holds configuration settings for the HEALTHCHECK feature.\nto see https://github.com/moby/moby/blob/master/api/types/container/config.go"
    },
    "mobyHostConfig": {
      "type": "object",
      "properties": {
        "auto_remove": {
          "type": "boolean",
          "format": "boolean"
        },
        "binds": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Applicable to all platforms"
        },
        "cap_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Applicable to UNIX platforms"
        },
        "cap_drop": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "cgroup": {
          "type": "string",
          "format": "string"
        },
        "console_size_height": {
          "type": "integer",
          "format": "int64",
          "title": "Applicable to Windows"
        },
        "console_size_width": {
          "type": "integer",
          "format": "int64"
        },
        "container_id_file": {
          "type": "string",
          "format": "string"
        },
        "dns": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "dns_options": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "dns_search": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "extra_hosts": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "group_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "init": {
          "type": "boolean",
          "format": "boolean",
          "title": "Run a custom init inside the container, if null, use the daemon's configured settings"
        },
        "ipc_mode": {
          "type": "string",
          "format": "string"
        },
        "isolation": {
          "type": "string",
          "format": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "log_config": {
          "$ref": "#/definitions/mobyLogConfig"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyVolumeMount"
          },
          "title": "Mounts specs used by the container"
        },
        "network_mode": {
          "type": "string",
          "format": "string"
        },
        "oom_score_adj": {
          "type": "integer",
          "format": "int32"
        },
        "pid_mode": {
          "type": "string",
          "format": "string"
        },
        "port_bindings": {
          "$ref": "#/definitions/mobyPortMap"
        },
        "privileged": {
          "type": "boolean",
          "format": "boolean"
        },
        "publish_all_ports": {
          "type": "boolean",
          "format": "boolean"
        },
        "readonly_rootfs": {
          "type": "boolean",
          "format": "boolean"
        },
        "resources": {
          "$ref": "#/definitions/mobyResources",
          "title": "Contains container's resources (cgroups, ulimits)"
        },
        "restart_policy": {
          "$ref": "#/definitions/mobyRestartPolicy"
        },
        "runtime": {
          "type": "string",
          "format": "string"
        },
        "security_opt": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "shm_size": {
          "type": "string",
          "format": "int64"
        },
        "storage_opt": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "sysctls": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "tmpfs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "userns_mode": {
          "type": "string",
          "format": "string"
        },
        "uts_mode": {
          "type": "string",
          "format": "string"
        },
        "volume_driver": {
          "type": "string",
          "format": "string"
        },
        "volumes_from": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "HostConfig the non-portable Config structure of a container.\nHere, \"non-portable\" means \"dependent of the host we are running on\".\nPortable information *should* appear in Config.\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyIPAM": {
      "type": "object",
      "properties": {
        "config": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyIPAMConfig"
          }
        },
        "driver": {
          "type": "string",
          "format": "string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "IPAM represents IP Address Management"
    },
    "mobyIPAMConfig": {
      "type": "object",
      "properties": {
        "aux_address": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "gateway": {
          "type": "string",
          "format": "string"
        },
        "ip_range": {
          "type": "string",
          "format": "string"
        },
        "subnet": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "IPAMConfig represents IPAM configurations"
    },
    "mobyLogConfig": {
      "type": "object",
      "properties": {
        "config": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "type": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "LogConfig represents the logging configuration of the container.\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyMountPoint": {
      "type": "object",
      "properties": {
        "destination": {
          "type": "string",
          "format": "string"
        },
        "driver": {
          "type": "string",
          "format": "string"
        },
        "mode": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "propagation": {
          "type": "string",
          "format": "string"
        },
        "rw": {
          "type": "boolean",
          "format": "boolean"
        },
        "source": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "MountPoint represents a mount point configuration inside the container."
    },
    "mobyNetworkCreate": {
      "type": "object",
      "properties": {
        "check_duplicate": {
          "type": "boolean",
          "format": "boolean"
        },
        "driver": {
          "type": "string",
          "format": "string"
        },
        "enable_ipv6": {
          "type": "boolean",
          "format": "boolean"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "ipam": {
          "$ref": "#/definitions/mobyIPAM"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "NetworkCreate is the expected body of the \"create network\" http request message"
    },
    "mobyNetworkCreateRequest": {
      "type": "object",
      "properties": {
        "Name": {
          "type": "string",
          "format": "string"
        },
        "network_create": {
          "$ref": "#/definitions/mobyNetworkCreate"
        }
      },
      "description": "NetworkCreateRequest is the request message sent to the server for network create call."
    },
    "mobyNetworkCreateResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string"
        },
        "warning": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "NetworkCreateResponse is the response message sent by the server for network create call"
    },
    "mobyNetworkResource": {
      "type": "object",
      "properties": {
        "containers": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointResource"
          }
        },
        "driver": {
          "type": "string",
          "format": "string"
        },
        "enable_ipv6": {
          "type": "boolean",
          "format": "boolean"
        },
        "id": {
          "type": "string",
          "format": "string"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean"
        },
        "ipam": {
          "$ref": "#/definitions/mobyIPAM"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "scope": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "NetworkResource is the body of the \"get network\" http response message"
    },
    "mobyNetworkSettings": {
      "type": "object",
      "properties": {
        "default_network_settings": {
          "$ref": "#/definitions/mobyDefaultNetworkSettings"
        },
        "network_settings_base": {
          "$ref": "#/definitions/mobyNetworkSettingsBase"
        },
        "networks": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointSettings"
          }
        }
      },
      "title": "NetworkSettings exposes the network settings in the api"
    },
    "mobyNetworkSettingsBase": {
      "type": "object",
      "properties": {
        "bridge": {
          "type": "string",
          "format": "string"
        },
        "hairpin_mode": {
          "type": "boolean",
          "format": "boolean"
        },
        "link_local_ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "link_local_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ports": {
          "$ref": "#/definitions/mobyPortMap"
        },
        "sandbox_id": {
          "type": "string",
          "format": "string"
        },
        "sandbox_key": {
          "type": "string",
          "format": "string"
        },
        "secondary_ip_addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyAddress"
          }
        },
        "secondary_ipv6_addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyAddress"
          }
        }
      },
      "title": "NetworkSettingsBase holds basic information about networks"
    },
    "mobyNetworkingConfig": {
      "type": "object",
      "properties": {
        "endpoints_config": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointSettings"
          }
        }
      },
      "title": "NetworkingConfig represents the container's networking configuration for each of its interfaces\nCarries the networking configs specified in the 'docker run' and 'docker network connect' commands\nto see https://github.com/moby/moby/blob/master/api/types/network/network.go"
    },
    "mobyPort": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string",
          "format": "string"
        },
        "private_port": {
          "type": "integer",
          "format": "int32"
        },
        "public_port": {
          "type": "integer",
          "format": "int32"
        },
        "type": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Port stores open ports info of container\ne.g. {\"PrivatePort\": 8080, \"PublicPort\": 80, \"Type\": \"tcp\"}"
    },
    "mobyPortBinding": {
      "type": "object",
      "properties": {
        "host_ip": {
          "type": "string",
          "format": "string",
          "title": "HostIP is the host IP Address"
        },
        "host_port": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "PortBinding represents a binding between a Host IP address and a Host Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyPortMap": {
      "type": "object",
      "properties": {
        "mapping_info": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/PortMapPortBindings"
          }
        },
        "value": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyPortBinding"
          }
        }
      },
      "title": "PortMap is a collection of PortBinding indexed by Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyPortSet": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "PortSet is a collection of structs indexed by Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyResources": {
      "type": "object",
      "properties": {
        "blkio_device_read_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyThrottleDevice"
          }
        },
        "blkio_device_read_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyThrottleDevice"
          }
        },
        "blkio_device_write_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyThrottleDevice"
          }
        },
        "blkio_device_write_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyThrottleDevice"
          }
        },
        "blkio_weight": {
          "type": "integer",
          "format": "int32"
        },
        "blkio_weight_device": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyWeightDevice"
          }
        },
        "cgroup_parent": {
          "type": "string",
          "format": "string",
          "title": "Applicable to UNIX platforms"
        },
        "cpu_count": {
          "type": "string",
          "format": "int64",
          "title": "Applicable to Windows"
        },
        "cpu_percent": {
          "type": "string",
          "format": "int64"
        },
        "cpu_period": {
          "type": "string",
          "format": "int64"
        },
        "cpu_quota": {
          "type": "string",
          "format": "int64"
        },
        "cpu_realtime_period": {
          "type": "string",
          "format": "int64"
        },
        "cpu_realtime_runtime": {
          "type": "string",
          "format": "int64"
        },
        "cpu_shares": {
          "type": "string",
          "format": "int64",
          "title": "Applicable to all platforms"
        },
        "cpuset_cpus": {
          "type": "string",
          "format": "string"
        },
        "cpuset_mems": {
          "type": "string",
          "format": "string"
        },
        "device_cgroup_rules": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "devices": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyDeviceMapping"
          }
        },
        "disk_quota": {
          "type": "string",
          "format": "int64"
        },
        "io_maximum_bandwidth": {
          "type": "string",
          "format": "uint64"
        },
        "io_maximum_iops": {
          "type": "string",
          "format": "uint64"
        },
        "kernel_memory": {
          "type": "string",
          "format": "int64"
        },
        "memory": {
          "type": "string",
          "format": "int64"
        },
        "memory_reservation": {
          "type": "string",
          "format": "int64"
        },
        "memory_swap": {
          "type": "string",
          "format": "int64"
        },
        "memory_swappiness": {
          "type": "string",
          "format": "int64"
        },
        "nano_cpus": {
          "type": "string",
          "format": "int64"
        },
        "oom_kill_disable": {
          "type": "boolean",
          "format": "boolean"
        },
        "pids_limit": {
          "type": "string",
          "format": "int64"
        },
        "ulimits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyUlimit"
          }
        }
      },
      "title": "Resources contains container's resources (cgroups config, ulimits...)\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyRestartPolicy": {
      "type": "object",
      "properties": {
        "maximum_retry_count": {
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "RestartPolicy represents the restart policies of the container."
    },
    "mobySummaryNetworkSettings": {
      "type": "object",
      "properties": {
        "networks": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointSettings"
          }
        }
      },
      "title": "SummaryNetworkSettings provides a summary of container's networks\nin /containers/json"
    },
    "mobyThrottleDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string",
          "format": "uint64"
        }
      },
      "title": "ThrottleDevice is a structure that holds device:rate_per_second pair\nto see http://github.com/moby/moby/blob/master/api/types/blkiodev/blkio.go"
    },
    "mobyTmpfsOptions": {
      "type": "object",
      "properties": {
        "mode": {
          "type": "integer",
          "format": "int64",
          "title": "Mode of the tmpfs upon creation"
        },
        "size_bytes": {
          "type": "string",
          "format": "int64",
          "description": "Size sets the size of the tmpfs, in bytes.\n\nThis will be converted to an operating system specific value\ndepending on the host. For example, on linux, it will be converted to\nuse a 'k', 'm' or 'g' syntax. BSD, though not widely supported with\ndocker, uses a straight byte value.\n\nPercentages are not supported."
        }
      },
      "description": "TmpfsOptions defines options specific to mounts of type \"tmpfs\"."
    },
    "mobyUlimit": {
      "type": "object",
      "properties": {
        "hard": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "soft": {
          "type": "string",
          "format": "int64"
        }
      },
      "title": "Ulimit is a human friendly version of Rlimit.\nto see https://github.com/docker/go-units/blob/master/ulimit.go"
    },
    "mobyVolumeMount": {
      "type": "object",
      "properties": {
        "bind_options": {
          "$ref": "#/definitions/mobyBindOptions"
        },
        "consistency": {
          "type": "string",
          "format": "string"
        },
        "mount_type": {
          "type": "string",
          "format": "string"
        },
        "read_only": {
          "type": "string",
          "format": "string"
        },
        "source": {
          "type": "string",
          "format": "string",
          "title": "Source specifies the name of the mount. Depending on mount type, this\nmay be a volume name or a host path, or even ignored.\nSource is not supported for tmpfs (must be an empty value)"
        },
        "target": {
          "type": "string",
          "format": "string"
        },
        "tmpfs_options": {
          "$ref": "#/definitions/mobyTmpfsOptions"
        },
        "volume_options": {
          "$ref": "#/definitions/mobyVolumeOptions"
        }
      },
      "title": "Mount represents a mount (volume).\nto see https://github.com/moby/moby/blob/master/api/types/mount/mount.go"
    },
    "mobyVolumeOptions": {
      "type": "object",
      "properties": {
        "driver": {
          "$ref": "#/definitions/mobyDriverConfig"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "no_copy": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "description": "VolumeOptions represents the options for a mount of type volume."
    },
    "mobyWeightDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string"
        },
        "weight": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "WeightDevice is a structure that holds device:weight pair\nto see http://github.com/moby/moby/blob/master/api/types/blkiodev/blkio.go"
    },
    "pbBridgedNetworkingData": {
      "type": "object",
      "properties": {
        "containers_networking": {
          "$ref": "#/definitions/BridgedNetworkingDataContainersNetworkingInfo"
        },
        "linux_bridges": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BridgedNetworkingDataLinuxBridgeInfo"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        },
        "veth_pairs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      }
    },
    "pbDockerNetworkCreationReqResp": {
      "type": "object",
      "properties": {
        "network_create_request": {
          "$ref": "#/definitions/mobyNetworkCreateRequest"
        },
        "network_create_response": {
          "$ref": "#/definitions/mobyNetworkCreateResponse"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDockerNetworkData": {
      "type": "object",
      "properties": {
        "network_resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyNetworkResource"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDockerProcessStatusReqResp": {
      "type": "object",
      "properties": {
        "containers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyContainer"
          }
        },
        "options": {
          "$ref": "#/definitions/mobyContainerListOptions"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDockerPullData": {
      "type": "object",
      "properties": {
        "image": {
          "type": "string",
          "format": "string"
        },
        "image_id": {
          "type": "string",
          "format": "string"
        },
        "progress_report": {
          "type": "string",
          "format": "string"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDockerRunData": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/mobyConfig"
        },
        "container_id": {
          "type": "string",
          "format": "string"
        },
        "container_name": {
          "type": "string",
          "format": "string"
        },
        "host_config": {
          "$ref": "#/definitions/mobyHostConfig"
        },
        "network_config": {
          "$ref": "#/definitions/mobyNetworkingConfig"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbEthernetSniffingData": {
      "type": "object",
      "properties": {
        "iface": {
          "type": "string",
          "format": "string"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        },
        "stats_and_packets": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      }
    },
    "pbInstantiationData": {
      "type": "object",
      "properties": {
        "instantiation": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyContainer"
          }
        },
        "metadata": {
          "$ref": "#/definitions/pbInstrumentMetadata"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "namespace": {
          "type": "string",
          "format": "string"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbInstrumentMetadata": {
      "type": "object",
      "properties": {
        "category_name": {
          "type": "string",
          "format": "string"
        },
        "class_name": {
          "type": "string",
          "format": "string"
        },
        "field_name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbProvisioningsData": {
      "type": "object",
      "properties": {
        "metadata": {
          "$ref": "#/definitions/ProvisioningsDataMetadata"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "namespace": {
          "type": "string",
          "format": "string"
        },
        "provisionings": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbDockerRunData"
          }
        }
      }
    },
    "pbRegistryRepositoryData": {
      "type": "object",
      "properties": {
        "registries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RegistryRepositoryDataRegistry"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
