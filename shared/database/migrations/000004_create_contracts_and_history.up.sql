CREATE TYPE contract_status AS ENUM ('draft', 'pending_signature', 'active', 'expired', 'cancelled');

CREATE TABLE contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    client_id UUID NOT NULL REFERENCES clients(id),
    title VARCHAR(255) NOT NULL,
    status contract_status NOT NULL DEFAULT 'draft',
    start_date DATE NOT NULL,
    end_date DATE,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
CREATE INDEX idx_contracts_tenant_deleted ON contracts(tenant_id, deleted_at);

CREATE TABLE contract_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id),
    contract_id UUID NOT NULL REFERENCES contracts(id),
    user_id UUID NOT NULL REFERENCES users(id),
    previous_status contract_status,
    new_status contract_status NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_history_contract ON contract_history(contract_id);
