/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package provider

type RediscloudSpec struct {
	// This is the Redis Cloud API key. It must be provided but can also be set by the `REDISCLOUD_ACCESS_KEY` environment variable.
	// +optional
	ApiKey *string `json:"apiKey,omitempty" tf:"api_key"`
	// This is the Redis Cloud API secret key. It must be provided but can also be set by the `REDISCLOUD_SECRET_KEY` environment variable.
	// +optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key"`
	// This is the URL of Redis Cloud and will default to `https://api.redislabs.com/v1`. This can also be set by the `REDISCLOUD_URL` environment variable.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}
