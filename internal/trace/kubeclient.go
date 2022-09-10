package trace

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeClient struct{ client.Client }

func NewKubeClient(c client.Client) *KubeClient {
	return &KubeClient{Client: c}
}

func (k *KubeClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.Get(ctx, key, obj, opts...)
}

func (k *KubeClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.List(ctx, list, opts...)
}

func (k *KubeClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.Create(ctx, obj, opts...)
}

func (k *KubeClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.Delete(ctx, obj, opts...)
}

func (k *KubeClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.Update(ctx, obj, opts...)
}

func (k *KubeClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	gvk := obj.GetObjectKind().GroupVersionKind()
	span.SetAttributes(attribute.String("debug.version", gvk.GroupVersion().String()))
	span.SetAttributes(attribute.String("debug.kind", gvk.Kind))
	span.SetAttributes(attribute.String("debug.name", obj.GetName()))
	span.SetAttributes(attribute.String("debug.namespace", obj.GetNamespace()))
	return k.Client.Patch(ctx, obj, patch, opts...)
}

func (k *KubeClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	span, ctx := StartSpanFromContext(ctx)
	defer span.End()
	return k.Client.DeleteAllOf(ctx, obj, opts...)
}
