package jsonCert

type JCert struct {
	Name             string      `json:"name"`
	MspId            string      `json:"mspid"`
	Roles            string      `json:"roles"`
	Affiliation      string      `json:"affiliation"`
	EnrollmentSecret string      `json:"enrollmentSecret"`
	Enrollment       JEnrollment `json:"enrollment"`
}

type JEnrollment struct {
	SigningIdentity string    `json:"signingIdentity"`
	Identity        JIdentity `json:"identity"`
}

type JIdentity struct {
	Certificate string `json:"certificate"`
}
