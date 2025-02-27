package sql

import (
	"context"

	"github.com/ory/kratos/corp"

	"github.com/gofrs/uuid"

	"github.com/ory/x/sqlcon"

	"github.com/ory/kratos/selfservice/flow/registration"
)

func (p *Persister) CreateRegistrationFlow(ctx context.Context, r *registration.Flow) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CreateRegistrationFlow")
	defer span.End()

	r.NID = corp.ContextualizeNID(ctx, p.nid)
	r.EnsureInternalContext()
	return p.GetConnection(ctx).Create(r)
}

func (p *Persister) UpdateRegistrationFlow(ctx context.Context, r *registration.Flow) error {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.UpdateRegistrationFlow")
	defer span.End()

	r.EnsureInternalContext()
	cp := *r
	cp.NID = corp.ContextualizeNID(ctx, p.nid)
	return p.update(ctx, cp)
}

func (p *Persister) GetRegistrationFlow(ctx context.Context, id uuid.UUID) (*registration.Flow, error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetRegistrationFlow")
	defer span.End()

	var r registration.Flow
	if err := p.GetConnection(ctx).Where("id = ? AND nid = ?",
		id, corp.ContextualizeNID(ctx, p.nid)).First(&r); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	return &r, nil
}
