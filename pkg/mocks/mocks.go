//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_core.go -source=../../capi/capi.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_controlplane.go -source=../../controlplane/controlplane.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_infrastructure.go -source=../../infrastructure/infrastructure.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_capa.go -source=../../infrastructure/capa/capa.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_kamaji.go -source=../../controlplane/kamaji/kamaji.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_kubeadm.go -source=../../bootstrap/kubeadm/kubeadm.go -build_flags=-mod=mod
//go:generate go run -mod=mod go.uber.org/mock/mockgen -package mocks -destination=./mock_sveltos.go -source=../../addon/sveltos/sveltos.go -build_flags=-mod=mod

package mocks

import _ "go.uber.org/mock/mockgen/model"
