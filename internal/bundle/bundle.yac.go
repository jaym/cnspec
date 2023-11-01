// Code generated by yac-it. DO NOT EDIT.
//
// Configure yac-it for things you want to auto-generate and extend generated
// objects in a separate file please.

package bundle

import (
	"encoding/json"
	"errors"
	"go.mondoo.com/cnquery/v9/explorer"
	"go.mondoo.com/cnspec/v9/policy"
	"gopkg.in/yaml.v3"
)

type FileContext struct {
	Line   int
	Column int
}

type Action explorer.Action

func (s *Action) UnmarshalYAML(node *yaml.Node) error {

	var decoded interface{}
	err := node.Decode(&decoded)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(decoded)
	if err != nil {
		return err
	}

	var v explorer.Action
	err = json.Unmarshal(jsonData, &v)
	if err == nil {
		*s = Action(v)
		return nil
	}

	return errors.New("failed to unmarshal Action")
}

type Author struct {
	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	Email       string      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" yaml:"email,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *Author) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Author
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Bundle struct {
	Packs         []*QueryPack    `protobuf:"bytes,10,rep,name=packs,proto3" json:"packs,omitempty" yaml:"packs,omitempty"`
	Frameworks    []*Framework    `protobuf:"bytes,8,rep,name=frameworks,proto3" json:"frameworks,omitempty" yaml:"frameworks,omitempty"`
	FrameworkMaps []*FrameworkMap `protobuf:"bytes,9,rep,name=framework_maps,json=frameworkMaps,proto3" json:"framework_maps,omitempty" yaml:"framework_maps,omitempty"`
	OwnerMrn      string          `protobuf:"bytes,1,opt,name=owner_mrn,json=ownerMrn,proto3" json:"owner_mrn,omitempty" yaml:"owner_mrn,omitempty"`
	Policies      []*Policy       `protobuf:"bytes,7,rep,name=policies,proto3" json:"policies,omitempty" yaml:"policies,omitempty"`
	Props         []*Property     `protobuf:"bytes,3,rep,name=props,proto3" json:"props,omitempty" yaml:"props,omitempty"`
	Queries       []*Mquery       `protobuf:"bytes,6,rep,name=queries,proto3" json:"queries,omitempty" yaml:"queries,omitempty"`
	Docs          *PolicyDocs     `protobuf:"bytes,5,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	FileContext   FileContext     `json:"-" yaml:"-"`
}

func (x *Bundle) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Bundle
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Control struct {
	Uid         string            `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string            `protobuf:"bytes,4,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Title       string            `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	Tags        map[string]string `protobuf:"bytes,34,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"tags,omitempty"`
	Docs        *ControlDocs      `protobuf:"bytes,21,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	Checksum    string            `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Action      Action            `protobuf:"varint,41,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	Manual      bool              `protobuf:"varint,50,opt,name=manual,proto3" json:"manual,omitempty" yaml:"manual,omitempty"`
	FileContext FileContext       `json:"-" yaml:"-"`
}

func (x *Control) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Control
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type ControlDocs struct {
	Refs        []*MqueryRef `protobuf:"bytes,4,rep,name=refs,proto3" json:"refs,omitempty" yaml:"refs,omitempty"`
	Desc        string       `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	FileContext FileContext  `json:"-" yaml:"-"`
}

func (x *ControlDocs) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp ControlDocs
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type ControlMap struct {
	Uid         string        `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string        `protobuf:"bytes,4,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Policies    []*ControlRef `protobuf:"bytes,8,rep,name=policies,proto3" json:"policies,omitempty" yaml:"policies,omitempty"`
	Checks      []*ControlRef `protobuf:"bytes,7,rep,name=checks,proto3" json:"checks,omitempty" yaml:"checks,omitempty"`
	Controls    []*ControlRef `protobuf:"bytes,9,rep,name=controls,proto3" json:"controls,omitempty" yaml:"controls,omitempty"`
	Queries     []*ControlRef `protobuf:"bytes,10,rep,name=queries,proto3" json:"queries,omitempty" yaml:"queries,omitempty"`
	FileContext FileContext   `json:"-" yaml:"-"`
}

func (x *ControlMap) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp ControlMap
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type ControlRef struct {
	Action      Action      `protobuf:"varint,41,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	Uid         string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string      `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *ControlRef) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp ControlRef
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Filters struct {
	Items       map[string]*Mquery `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"items,omitempty"`
	FileContext FileContext        `json:"-" yaml:"-"`
}

func (x *Filters) addFileContext(node *yaml.Node) {
	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
}

type Framework struct {
	Summary                string            `protobuf:"bytes,46,opt,name=summary,proto3" json:"summary,omitempty" yaml:"summary,omitempty"`
	Created                int64             `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified               int64             `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	Dependencies           []*FrameworkRef   `protobuf:"bytes,35,rep,name=dependencies,proto3" json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	LocalContentChecksum   string            `protobuf:"bytes,37,opt,name=local_content_checksum,json=localContentChecksum,proto3" json:"local_content_checksum,omitempty" yaml:"local_content_checksum,omitempty"`
	GraphContentChecksum   string            `protobuf:"bytes,38,opt,name=graph_content_checksum,json=graphContentChecksum,proto3" json:"graph_content_checksum,omitempty" yaml:"graph_content_checksum,omitempty"`
	LocalExecutionChecksum string            `protobuf:"bytes,39,opt,name=local_execution_checksum,json=localExecutionChecksum,proto3" json:"local_execution_checksum,omitempty" yaml:"local_execution_checksum,omitempty"`
	GraphExecutionChecksum string            `protobuf:"bytes,40,opt,name=graph_execution_checksum,json=graphExecutionChecksum,proto3" json:"graph_execution_checksum,omitempty" yaml:"graph_execution_checksum,omitempty"`
	FrameworkMaps          []*FrameworkMap   `protobuf:"bytes,53,rep,name=framework_maps,json=frameworkMaps,proto3" json:"framework_maps,omitempty" yaml:"framework_maps,omitempty"`
	Uid                    string            `protobuf:"bytes,36,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	OwnerMrn               string            `protobuf:"bytes,8,opt,name=owner_mrn,json=ownerMrn,proto3" json:"owner_mrn,omitempty" yaml:"owner_mrn,omitempty"`
	Mrn                    string            `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Name                   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	Version                string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" yaml:"version,omitempty"`
	License                string            `protobuf:"bytes,21,opt,name=license,proto3" json:"license,omitempty" yaml:"license,omitempty"`
	Tags                   map[string]string `protobuf:"bytes,34,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"tags,omitempty"`
	Authors                []*Author         `protobuf:"bytes,30,rep,name=authors,proto3" json:"authors,omitempty" yaml:"authors,omitempty"`
	Docs                   *PolicyDocs       `protobuf:"bytes,41,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	Groups                 []*FrameworkGroup `protobuf:"bytes,11,rep,name=groups,proto3" json:"groups,omitempty" yaml:"groups,omitempty"`
	FileContext            FileContext       `json:"-" yaml:"-"`
}

func (x *Framework) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Framework
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type FrameworkGroup struct {
	Uid          string           `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Title        string           `protobuf:"bytes,24,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	Type         GroupType        `protobuf:"varint,4,opt,name=type,proto3,enum=cnspec.policy.v1.GroupType" json:"type,omitempty" yaml:"type,omitempty"`
	StartDate    int64            `protobuf:"varint,21,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty" yaml:"start_date,omitempty"`
	EndDate      int64            `protobuf:"varint,22,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty" yaml:"end_date,omitempty"`
	ReminderDate int64            `protobuf:"varint,23,opt,name=reminder_date,json=reminderDate,proto3" json:"reminder_date,omitempty" yaml:"reminder_date,omitempty"`
	Reviewers    []*Author        `protobuf:"bytes,27,rep,name=reviewers,proto3" json:"reviewers,omitempty" yaml:"reviewers,omitempty"`
	ReviewStatus ReviewStatus     `protobuf:"varint,28,opt,name=review_status,json=reviewStatus,proto3,enum=cnspec.policy.v1.ReviewStatus" json:"review_status,omitempty" yaml:"review_status,omitempty"`
	Created      int64            `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified     int64            `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	Controls     []*Control       `protobuf:"bytes,1,rep,name=controls,proto3" json:"controls,omitempty" yaml:"controls,omitempty"`
	Authors      []*Author        `protobuf:"bytes,26,rep,name=authors,proto3" json:"authors,omitempty" yaml:"authors,omitempty"`
	Docs         *PolicyGroupDocs `protobuf:"bytes,25,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	FileContext  FileContext      `json:"-" yaml:"-"`
}

func (x *FrameworkGroup) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp FrameworkGroup
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type FrameworkMap struct {
	PolicyDependencies     []*ObjectRef  `protobuf:"bytes,4,rep,name=policy_dependencies,json=policyDependencies,proto3" json:"policy_dependencies,omitempty" yaml:"policy_dependencies,omitempty"`
	QueryPackDependencies  []*ObjectRef  `protobuf:"bytes,6,rep,name=query_pack_dependencies,json=queryPackDependencies,proto3" json:"query_pack_dependencies,omitempty" yaml:"query_pack_dependencies,omitempty"`
	LocalContentChecksum   string        `protobuf:"bytes,21,opt,name=local_content_checksum,json=localContentChecksum,proto3" json:"local_content_checksum,omitempty" yaml:"local_content_checksum,omitempty"`
	LocalExecutionChecksum string        `protobuf:"bytes,22,opt,name=local_execution_checksum,json=localExecutionChecksum,proto3" json:"local_execution_checksum,omitempty" yaml:"local_execution_checksum,omitempty"`
	Uid                    string        `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn                    string        `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	FrameworkOwner         *ObjectRef    `protobuf:"bytes,20,opt,name=framework_owner,json=frameworkOwner,proto3" json:"framework_owner,omitempty" yaml:"framework_owner,omitempty"`
	FrameworkDependencies  []*ObjectRef  `protobuf:"bytes,3,rep,name=framework_dependencies,json=frameworkDependencies,proto3" json:"framework_dependencies,omitempty" yaml:"framework_dependencies,omitempty"`
	Controls               []*ControlMap `protobuf:"bytes,5,rep,name=controls,proto3" json:"controls,omitempty" yaml:"controls,omitempty"`
	FileContext            FileContext   `json:"-" yaml:"-"`
}

func (x *FrameworkMap) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp FrameworkMap
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type FrameworkRef struct {
	Action      Action      `protobuf:"varint,41,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	Uid         string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string      `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *FrameworkRef) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp FrameworkRef
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type GroupType policy.GroupType

func (s *GroupType) UnmarshalYAML(node *yaml.Node) error {

	var decoded interface{}
	err := node.Decode(&decoded)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(decoded)
	if err != nil {
		return err
	}

	var v policy.GroupType
	err = json.Unmarshal(jsonData, &v)
	if err == nil {
		*s = GroupType(v)
		return nil
	}

	return errors.New("failed to unmarshal GroupType")
}

type Impact struct {
	Value       *ImpactValue           `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" yaml:"value,omitempty"`
	Scoring     explorer.ScoringSystem `protobuf:"varint,2,opt,name=scoring,proto3,enum=cnquery.explorer.ScoringSystem" json:"scoring,omitempty" yaml:"scoring,omitempty"`
	Weight      int32                  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty" yaml:"weight,omitempty"`
	Action      Action                 `protobuf:"varint,4,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	FileContext FileContext            `json:"-" yaml:"-"`
}

func (x *Impact) addFileContext(node *yaml.Node) {
	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
}

type ImpactValue struct {
	Value       int32       `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty" yaml:"value,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *ImpactValue) addFileContext(node *yaml.Node) {
	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
}

type Mquery struct {
	Uid         string            `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string            `protobuf:"bytes,4,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Title       string            `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	Impact      *Impact           `protobuf:"bytes,23,opt,name=impact,proto3" json:"impact,omitempty" yaml:"impact,omitempty"`
	Filters     *Filters          `protobuf:"bytes,37,opt,name=filters,proto3" json:"filters,omitempty" yaml:"filters,omitempty"`
	Tags        map[string]string `protobuf:"bytes,34,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"tags,omitempty"`
	Props       []*Property       `protobuf:"bytes,38,rep,name=props,proto3" json:"props,omitempty" yaml:"props,omitempty"`
	Mql         string            `protobuf:"bytes,1,opt,name=mql,proto3" json:"mql,omitempty" yaml:"mql,omitempty"`
	Docs        *MqueryDocs       `protobuf:"bytes,21,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	Refs        []*MqueryRef      `protobuf:"bytes,22,rep,name=refs,proto3" json:"refs,omitempty" yaml:"refs,omitempty"`
	Query       string            `protobuf:"bytes,40,opt,name=query,proto3" json:"query,omitempty" yaml:"query,omitempty"`
	CodeId      string            `protobuf:"bytes,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id,omitempty"`
	Checksum    string            `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Type        string            `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty" yaml:"type,omitempty"`
	Context     string            `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty" yaml:"context,omitempty"`
	Desc        string            `protobuf:"bytes,35,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	Variants    []*ObjectRef      `protobuf:"bytes,39,rep,name=variants,proto3" json:"variants,omitempty" yaml:"variants,omitempty"`
	Action      Action            `protobuf:"varint,41,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	FileContext FileContext       `json:"-" yaml:"-"`
}

func (x *Mquery) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Mquery
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type MqueryDocs struct {
	Refs        []*MqueryRef `protobuf:"bytes,4,rep,name=refs,proto3" json:"refs,omitempty" yaml:"refs,omitempty"`
	Desc        string       `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	Audit       string       `protobuf:"bytes,2,opt,name=audit,proto3" json:"audit,omitempty" yaml:"audit,omitempty"`
	Remediation *Remediation `protobuf:"bytes,5,opt,name=remediation,proto3" json:"remediation,omitempty" yaml:"remediation,omitempty"`
	FileContext FileContext  `json:"-" yaml:"-"`
}

func (x *MqueryDocs) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp MqueryDocs
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type MqueryRef struct {
	Url         string      `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" yaml:"url,omitempty"`
	Title       string      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *MqueryRef) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp MqueryRef
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type ObjectRef struct {
	Uid         string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string      `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *ObjectRef) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp ObjectRef
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Policy struct {
	Summary                string                 `protobuf:"bytes,46,opt,name=summary,proto3" json:"summary,omitempty" yaml:"summary,omitempty"`
	Created                int64                  `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified               int64                  `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	LocalContentChecksum   string                 `protobuf:"bytes,37,opt,name=local_content_checksum,json=localContentChecksum,proto3" json:"local_content_checksum,omitempty" yaml:"local_content_checksum,omitempty"`
	GraphContentChecksum   string                 `protobuf:"bytes,38,opt,name=graph_content_checksum,json=graphContentChecksum,proto3" json:"graph_content_checksum,omitempty" yaml:"graph_content_checksum,omitempty"`
	LocalExecutionChecksum string                 `protobuf:"bytes,39,opt,name=local_execution_checksum,json=localExecutionChecksum,proto3" json:"local_execution_checksum,omitempty" yaml:"local_execution_checksum,omitempty"`
	GraphExecutionChecksum string                 `protobuf:"bytes,40,opt,name=graph_execution_checksum,json=graphExecutionChecksum,proto3" json:"graph_execution_checksum,omitempty" yaml:"graph_execution_checksum,omitempty"`
	ComputedFilters        *Filters               `protobuf:"bytes,43,opt,name=computed_filters,json=computedFilters,proto3" json:"computed_filters,omitempty" yaml:"computed_filters,omitempty"`
	QueryCounts            *QueryCounts           `protobuf:"bytes,42,opt,name=query_counts,json=queryCounts,proto3" json:"query_counts,omitempty" yaml:"query_counts,omitempty"`
	Uid                    string                 `protobuf:"bytes,36,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	OwnerMrn               string                 `protobuf:"bytes,8,opt,name=owner_mrn,json=ownerMrn,proto3" json:"owner_mrn,omitempty" yaml:"owner_mrn,omitempty"`
	Mrn                    string                 `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Name                   string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	Version                string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" yaml:"version,omitempty"`
	License                string                 `protobuf:"bytes,21,opt,name=license,proto3" json:"license,omitempty" yaml:"license,omitempty"`
	Tags                   map[string]string      `protobuf:"bytes,34,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"tags,omitempty"`
	Props                  []*Property            `protobuf:"bytes,45,rep,name=props,proto3" json:"props,omitempty" yaml:"props,omitempty"`
	Authors                []*Author              `protobuf:"bytes,30,rep,name=authors,proto3" json:"authors,omitempty" yaml:"authors,omitempty"`
	Docs                   *PolicyDocs            `protobuf:"bytes,41,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	Groups                 []*PolicyGroup         `protobuf:"bytes,11,rep,name=groups,proto3" json:"groups,omitempty" yaml:"groups,omitempty"`
	ScoringSystem          explorer.ScoringSystem `protobuf:"varint,10,opt,name=scoring_system,json=scoringSystem,proto3,enum=cnquery.explorer.ScoringSystem" json:"scoring_system,omitempty" yaml:"scoring_system,omitempty"`
	FileContext            FileContext            `json:"-" yaml:"-"`
}

func (x *Policy) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Policy
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type PolicyDocs struct {
	Desc        string      `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *PolicyDocs) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp PolicyDocs
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type PolicyGroup struct {
	Uid          string           `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Type         GroupType        `protobuf:"varint,4,opt,name=type,proto3,enum=cnspec.policy.v1.GroupType" json:"type,omitempty" yaml:"type,omitempty"`
	StartDate    int64            `protobuf:"varint,21,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty" yaml:"start_date,omitempty"`
	EndDate      int64            `protobuf:"varint,22,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty" yaml:"end_date,omitempty"`
	ReminderDate int64            `protobuf:"varint,23,opt,name=reminder_date,json=reminderDate,proto3" json:"reminder_date,omitempty" yaml:"reminder_date,omitempty"`
	Created      int64            `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified     int64            `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	Policies     []*PolicyRef     `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty" yaml:"policies,omitempty"`
	Title        string           `protobuf:"bytes,24,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	Filters      *Filters         `protobuf:"bytes,20,opt,name=filters,proto3" json:"filters,omitempty" yaml:"filters,omitempty"`
	Checks       []*Mquery        `protobuf:"bytes,2,rep,name=checks,proto3" json:"checks,omitempty" yaml:"checks,omitempty"`
	Queries      []*Mquery        `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty" yaml:"queries,omitempty"`
	Authors      []*Author        `protobuf:"bytes,26,rep,name=authors,proto3" json:"authors,omitempty" yaml:"authors,omitempty"`
	Docs         *PolicyGroupDocs `protobuf:"bytes,25,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	FileContext  FileContext      `json:"-" yaml:"-"`
}

func (x *PolicyGroup) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp PolicyGroup
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type PolicyGroupDocs struct {
	Desc          string      `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	Justification string      `protobuf:"bytes,2,opt,name=justification,proto3" json:"justification,omitempty" yaml:"justification,omitempty"`
	FileContext   FileContext `json:"-" yaml:"-"`
}

func (x *PolicyGroupDocs) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp PolicyGroupDocs
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type PolicyRef struct {
	Action      Action      `protobuf:"varint,41,opt,name=action,proto3,enum=cnquery.explorer.Action" json:"action,omitempty" yaml:"action,omitempty"`
	Checksum    string      `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Uid         string      `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string      `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Impact      *Impact     `protobuf:"bytes,23,opt,name=impact,proto3" json:"impact,omitempty" yaml:"impact,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *PolicyRef) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp PolicyRef
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Property struct {
	Uid         string       `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	Mrn         string       `protobuf:"bytes,4,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Title       string       `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	CodeId      string       `protobuf:"bytes,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id,omitempty"`
	Checksum    string       `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Type        string       `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty" yaml:"type,omitempty"`
	Context     string       `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty" yaml:"context,omitempty"`
	For         []*ObjectRef `protobuf:"bytes,8,rep,name=for,proto3" json:"for,omitempty" yaml:"for,omitempty"`
	Desc        string       `protobuf:"bytes,35,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	Mql         string       `protobuf:"bytes,1,opt,name=mql,proto3" json:"mql,omitempty" yaml:"mql,omitempty"`
	FileContext FileContext  `json:"-" yaml:"-"`
}

func (x *Property) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp Property
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type QueryCounts struct {
	ScoringCount int64       `protobuf:"varint,1,opt,name=scoring_count,json=scoringCount,proto3" json:"scoring_count,omitempty" yaml:"scoring_count,omitempty"`
	DataCount    int64       `protobuf:"varint,2,opt,name=data_count,json=dataCount,proto3" json:"data_count,omitempty" yaml:"data_count,omitempty"`
	TotalCount   int64       `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty" yaml:"total_count,omitempty"`
	FileContext  FileContext `json:"-" yaml:"-"`
}

func (x *QueryCounts) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp QueryCounts
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type QueryGroup struct {
	Created     int64       `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified    int64       `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	Title       string      `protobuf:"bytes,24,opt,name=title,proto3" json:"title,omitempty" yaml:"title,omitempty"`
	Filters     *Filters    `protobuf:"bytes,20,opt,name=filters,proto3" json:"filters,omitempty" yaml:"filters,omitempty"`
	Queries     []*Mquery   `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty" yaml:"queries,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *QueryGroup) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp QueryGroup
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type QueryPack struct {
	LocalContentChecksum   string             `protobuf:"bytes,23,opt,name=local_content_checksum,json=localContentChecksum,proto3" json:"local_content_checksum,omitempty" yaml:"local_content_checksum,omitempty"`
	LocalExecutionChecksum string             `protobuf:"bytes,24,opt,name=local_execution_checksum,json=localExecutionChecksum,proto3" json:"local_execution_checksum,omitempty" yaml:"local_execution_checksum,omitempty"`
	Uid                    string             `protobuf:"bytes,36,opt,name=uid,proto3" json:"uid,omitempty" yaml:"uid,omitempty"`
	OwnerMrn               string             `protobuf:"bytes,4,opt,name=owner_mrn,json=ownerMrn,proto3" json:"owner_mrn,omitempty" yaml:"owner_mrn,omitempty"`
	Mrn                    string             `protobuf:"bytes,1,opt,name=mrn,proto3" json:"mrn,omitempty" yaml:"mrn,omitempty"`
	Name                   string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	Version                string             `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" yaml:"version,omitempty"`
	License                string             `protobuf:"bytes,21,opt,name=license,proto3" json:"license,omitempty" yaml:"license,omitempty"`
	Filters                *Filters           `protobuf:"bytes,48,opt,name=filters,proto3" json:"filters,omitempty" yaml:"filters,omitempty"`
	Props                  []*Property        `protobuf:"bytes,35,rep,name=props,proto3" json:"props,omitempty" yaml:"props,omitempty"`
	Tags                   map[string]string  `protobuf:"bytes,34,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"tags,omitempty"`
	Queries                []*Mquery          `protobuf:"bytes,6,rep,name=queries,proto3" json:"queries,omitempty" yaml:"queries,omitempty"`
	Authors                []*Author          `protobuf:"bytes,30,rep,name=authors,proto3" json:"authors,omitempty" yaml:"authors,omitempty"`
	Docs                   *QueryPackDocs     `protobuf:"bytes,22,opt,name=docs,proto3" json:"docs,omitempty" yaml:"docs,omitempty"`
	Groups                 []*QueryGroup      `protobuf:"bytes,11,rep,name=groups,proto3" json:"groups,omitempty" yaml:"groups,omitempty"`
	AssetFilters           map[string]*Mquery `protobuf:"bytes,7,rep,name=asset_filters,json=assetFilters,proto3" json:"asset_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" yaml:"asset_filters,omitempty"`
	DeprecatedFilters      []string           `protobuf:"bytes,43,rep,name=deprecated_filters,json=deprecatedFilters,proto3" json:"deprecated_filters,omitempty" yaml:"deprecated_filters,omitempty"`
	Context                string             `protobuf:"bytes,8,opt,name=context,proto3" json:"context,omitempty" yaml:"context,omitempty"`
	ComputedFilters        *Filters           `protobuf:"bytes,47,opt,name=computed_filters,json=computedFilters,proto3" json:"computed_filters,omitempty" yaml:"computed_filters,omitempty"`
	Summary                string             `protobuf:"bytes,46,opt,name=summary,proto3" json:"summary,omitempty" yaml:"summary,omitempty"`
	Created                int64              `protobuf:"varint,32,opt,name=created,proto3" json:"created,omitempty" yaml:"created,omitempty"`
	Modified               int64              `protobuf:"varint,33,opt,name=modified,proto3" json:"modified,omitempty" yaml:"modified,omitempty"`
	FileContext            FileContext        `json:"-" yaml:"-"`
}

func (x *QueryPack) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp QueryPack
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type QueryPackDocs struct {
	Desc        string      `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *QueryPackDocs) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp QueryPackDocs
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}

type Remediation struct {
	Items       []*TypedDoc `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" yaml:"items,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *Remediation) addFileContext(node *yaml.Node) {
	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
}

type ReviewStatus policy.ReviewStatus

func (s *ReviewStatus) UnmarshalYAML(node *yaml.Node) error {

	var decoded interface{}
	err := node.Decode(&decoded)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(decoded)
	if err != nil {
		return err
	}

	var v policy.ReviewStatus
	err = json.Unmarshal(jsonData, &v)
	if err == nil {
		*s = ReviewStatus(v)
		return nil
	}

	return errors.New("failed to unmarshal ReviewStatus")
}

type TypedDoc struct {
	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id,omitempty"`
	Desc        string      `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc,omitempty"`
	FileContext FileContext `json:"-" yaml:"-"`
}

func (x *TypedDoc) UnmarshalYAML(node *yaml.Node) error {
	// prevent recursive calls into UnmarshalYAML with a placeholder type
	type tmp TypedDoc
	err := node.Decode((*tmp)(x))
	if err != nil {
		return err
	}

	x.FileContext.Column = node.Column
	x.FileContext.Line = node.Line
	return nil
}
