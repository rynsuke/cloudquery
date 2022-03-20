-- Autogenerated by migration tool on 2022-03-14 15:54:35

-- Resource: ec2.eips
ALTER TABLE IF EXISTS aws_ec2_eips
DROP CONSTRAINT aws_ec2_eips_pk;
ALTER TABLE IF EXISTS aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (cq_fetch_date, account_id, allocation_id);




-- Resource: cloudformation.stacks
CREATE TABLE IF NOT EXISTS "aws_cloudformation_stacks" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"arn" text,
	"creation_time" timestamp without time zone,
	"stack" text,
	"status" text,
	"capabilities" text[],
	"change_set_id" text,
	"deletion_time" timestamp without time zone,
	"description" text,
	"disable_rollback" boolean,
	"stack_drift_status" text,
	"drift_last_check_timestamp" timestamp without time zone,
	"enable_termination_protection" boolean,
	"last_updated_time" timestamp without time zone,
	"notification_arns" text[],
	"parameters" jsonb,
	"parent_id" text,
	"role_arn" text,
	"rollback_configuration_monitoring_time_in_minutes" integer,
	"rollback_configuration_rollback_triggers" jsonb,
	"root_id" text,
	"id" text,
	"stack_status_reason" text,
	"tags" jsonb,
	"timeout_in_minutes" integer,
	CONSTRAINT aws_cloudformation_stacks_pk PRIMARY KEY(cq_fetch_date,id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_cloudformation_stacks');
CREATE TABLE IF NOT EXISTS "aws_cloudformation_stack_outputs" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"stack_cq_id" uuid,
	"description" text,
	"export_name" text,
	"output_key" text,
	"output_value" text,
	CONSTRAINT aws_cloudformation_stack_outputs_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_cloudformation_stack_outputs (cq_fetch_date, stack_cq_id);
SELECT setup_tsdb_child('aws_cloudformation_stack_outputs', 'stack_cq_id', 'aws_cloudformation_stacks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_cloudformation_stack_resources" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"stack_cq_id" uuid,
	"last_updated_timestamp" timestamp without time zone,
	"logical_resource_id" text,
	"resource_status" text,
	"resource_type" text,
	"stack_resource_drift_status" text,
	"drift_last_check_timestamp" timestamp without time zone,
	"module_info_logical_id_hierarchy" text,
	"module_info_type_hierarchy" text,
	"physical_resource_id" text,
	"resource_status_reason" text,
	CONSTRAINT aws_cloudformation_stack_resources_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_cloudformation_stack_resources (cq_fetch_date, stack_cq_id);
SELECT setup_tsdb_child('aws_cloudformation_stack_resources', 'stack_cq_id', 'aws_cloudformation_stacks', 'cq_id');
