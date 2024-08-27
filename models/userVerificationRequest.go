package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserVerificationRequest struct {
	ID                        primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	Case                      int                    `bson:"case,omitempty" json:"case,omitempty"`
	BCAId                     *primitive.ObjectID    `bson:"bcaId,omitempty" json:"bcaId,omitempty"`
	VendorId                  *primitive.ObjectID    `bson:"vendorId,omitempty" json:"vendorId,omitempty"`
	CompanyId                 primitive.ObjectID     `bson:"companyId" json:"companyId"`
	UserInfo                  map[string]interface{} `bson:"userInfo,omitempty" json:"userInfo,omitempty"`
	BCAInfo                   map[string]interface{} `bson:"bcaInfo,omitempty" json:"bcaInfo,omitempty"`
	CompanyInfo               map[string]interface{} `bson:"companyInfo,omitempty" json:"companyInfo,omitempty"`
	Status                    string                 `bson:"status,omitempty" json:"status,omitempty"`
	Priority                  string                 `bson:"priority,omitempty" json:"priority,omitempty"`
	LinkStatus                string                 `bson:"linkStatus,omitempty" json:"linkStatus,omitempty"`
	Hash                      *string                `bson:"hash,omitempty" json:"hash,omitempty"`
	CoordinatorName           *string                `bson:"coordinatorName,omitempty" json:"coordinatorName,omitempty"`
	Checks                    []string               `bson:"checks" json:"checks"`
	ReportDetail              map[string]interface{} `bson:"reportDetail,omitempty" json:"reportDetail,omitempty"`
	FinalReportInfo           map[string]interface{} `bson:"finalReportInfo,omitempty" json:"finalReportInfo,omitempty"`
	IsAssigned                bool                   `bson:"isAssigned" json:"isAssigned"`
	IsDetailFilled            bool                   `bson:"isDetailFilled" json:"isDetailFilled"`
	IsInefficiency            bool                   `bson:"isInefficiency" json:"isInefficiency"`
	EducationTaskId           *primitive.ObjectID    `bson:"educationTaskId,omitempty" json:"educationTaskId,omitempty"`
	ExperienceTaskId          *primitive.ObjectID    `bson:"experienceTaskId,omitempty" json:"experienceTaskId,omitempty"`
	CertificateTaskId         *primitive.ObjectID    `bson:"certificateTaskId,omitempty" json:"certificateTaskId,omitempty"`
	IdentityTaskId            *primitive.ObjectID    `bson:"identityTaskId,omitempty" json:"identityTaskId,omitempty"`
	PoliceTaskId              *primitive.ObjectID    `bson:"policeTaskId,omitempty" json:"policeTaskId,omitempty"`
	CourtTaskId               *primitive.ObjectID    `bson:"courtTaskId,omitempty" json:"courtTaskId,omitempty"`
	AddressTaskId             *primitive.ObjectID    `bson:"addressTaskId,omitempty" json:"addressTaskId,omitempty"`
	DiscreteCallsTaskId       *primitive.ObjectID    `bson:"discreteCallsTaskId,omitempty" json:"discreteCallsTaskId,omitempty"`
	VoloHealthCareTaskId      *primitive.ObjectID    `bson:"voloHealthCareTaskId,omitempty" json:"voloHealthCareTaskId,omitempty"`
	KEIProcessTaskId          *primitive.ObjectID    `bson:"keiProcessTaskId,omitempty" json:"keiProcessTaskId,omitempty"`
	EnhanceDueDiligenceTaskId *primitive.ObjectID    `bson:"enhanceDueDiligenceTaskId,omitempty" json:"enhanceDueDiligenceTaskId,omitempty"`
	REEDDAppointmentTaskId    *primitive.ObjectID    `bson:"reeddAppointmentTaskId,omitempty" json:"reeddAppointmentTaskId,omitempty"`
	REEDDSurpriseTaskId       *primitive.ObjectID    `bson:"reeddSurpriseTaskId,omitempty" json:"reeddSurpriseTaskId,omitempty"`
	DrugTaskId                *primitive.ObjectID    `bson:"drugTaskId,omitempty" json:"drugTaskId,omitempty"`
	CreditTaskId              *primitive.ObjectID    `bson:"creditTaskId,omitempty" json:"creditTaskId,omitempty"`
	ReferenceTaskId           *primitive.ObjectID    `bson:"referenceTaskId,omitempty" json:"referenceTaskId,omitempty"`
	VideoKycTaskId            *primitive.ObjectID    `bson:"videoKycTaskId,omitempty" json:"videoKycTaskId,omitempty"`
	GlobalDatabaseTaskId      *primitive.ObjectID    `bson:"globalDatabaseTaskId,omitempty" json:"globalDatabaseTaskId,omitempty"`
	SexOffenderTaskId         *primitive.ObjectID    `bson:"sexOffenderTaskId,omitempty" json:"sexOffenderTaskId,omitempty"`
	TransactionId             *primitive.ObjectID    `bson:"transactionId,omitempty" json:"transactionId,omitempty"`
	TransactionStatus         *string                `bson:"transactionStatus,omitempty" json:"transactionStatus,omitempty"`
	RazorpayInvoiceId         *string                `bson:"razorpayInvoiceId,omitempty" json:"razorpayInvoiceId,omitempty"`
	RazorpayInvoiceUrl        *string                `bson:"razorpayInvoiceUrl,omitempty" json:"razorpayInvoiceUrl,omitempty"`
	RequestedDate             time.Time              `bson:"requestedDate,omitempty" json:"requestedDate,omitempty"`
	DeadlineDate              time.Time              `bson:"deadlineDate,omitempty" json:"deadlineDate,omitempty"`
	IsREEDDSURPRISE           bool                   `bson:"isREEDDSURPRISE" json:"isREEDDSURPRISE"`
	IsREEDDAPPOINTMENT        bool                   `bson:"isREEDDAPPOINTMENT" json:"isREEDDAPPOINTMENT"`
	IsEDD                     bool                   `bson:"isEDD" json:"isEDD"`
	IsDISCRETE                bool                   `bson:"isDISCRETE" json:"isDISCRETE"`
	IsKEI                     bool                   `bson:"isKEI" json:"isKEI"`
	IsVOLO                    bool                   `bson:"isVOLO" json:"isVOLO"`
	FinalReportDate           *time.Time             `bson:"finalReportDate,omitempty" json:"finalReportDate,omitempty"`
	ReferredByCompany         bool                   `bson:"referredByCompany" json:"referredByCompany"`
	IsDeleted                 bool                   `bson:"isDeleted" json:"isDeleted"`
	CreatedAt                 time.Time              `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt                 time.Time              `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

// BeforeSave handles the logic before saving a new record
func (u *UserVerificationRequest) BeforeSave(ctx context.Context, db *mongo.Database) error {
	counter := &Counter{}
	err := db.Collection("counters").FindOneAndUpdate(
		ctx,
		bson.M{"_id": "userVerificationRequest"},
		bson.M{"$inc": bson.M{"case": 1}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	).Decode(counter)
	if err != nil {
		return err
	}
	u.Case = counter.Case
	return nil
}

type Counter struct {
	ID   string `bson:"_id"`
	Case int    `bson:"case"`
}
