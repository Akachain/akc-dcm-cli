package jsonCert

type JCert struct {
	Name             string      `jsonCert:"name"`
	MspId            string      `jsonCert:"mspid"`
	Roles            string      `jsonCert:"roles"`
	Affiliation      string      `jsonCert:"affiliation"`
	EnrollmentSecret string      `jsonCert:"enrollmentSecret"`
	Enrollment       JEnrollment `jsonCert:"enrollment"`
}

type JEnrollment struct {
	SigningIdentity string    `jsonCert:"signingIdentity"`
	Identity        JIdentity `jsonCert:"identity"`
}

type JIdentity struct {
	Certificate string `jsonCert:"certificate"`
}
