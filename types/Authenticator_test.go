package types

import (
	"encoding/hex"
	"fmt"
	"github.com/jcmturner/gokrb5/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func unmarshalAuthenticator_test(t *testing.T, v string) Authenticator {
	var a Authenticator
	//t.Logf("Starting unmarshal tests of %s", v)
	b, err := hex.DecodeString(testdata.TestVectors[v])
	if err != nil {
		t.Fatalf("Test vector read error of %s: %v\n", v, err)
	}
	err = a.Unmarshal(b)
	if err != nil {
		t.Fatalf("Unmarshal error of %s: %v\n", v, err)
	}
	return a
}
func TestUnmarshalAuthenticator(t *testing.T) {
	a := unmarshalAuthenticator_test(t, "encode_krb5_authenticator")
	//Parse the test time value into a time.Time type
	tt, _ := time.Parse(testdata.TEST_TIME_FORMAT, testdata.TEST_TIME)

	assert.Equal(t, testdata.TEST_KVNO, a.AVNO, "Authenticator version number not as expected")
	assert.Equal(t, testdata.TEST_REALM, a.CRealm, "CRealm not as expected")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMETYPE, a.CName.NameType, "CName NameType not as expected")
	assert.Equal(t, len(testdata.TEST_PRINCIPALNAME_NAMESTRING), len(a.CName.NameString), "CName does not have the expected number of NameStrings")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMESTRING, a.CName.NameString, "CName entries not as expected")
	assert.Equal(t, 1, a.Cksum.CksumType, "Checksum type not as expected")
	assert.Equal(t, []byte("1234"), a.Cksum.Checksum, "Checsum not as expected")
	assert.Equal(t, 123456, a.Cusec, "Client microseconds not as expected")
	assert.Equal(t, tt, a.CTime, "Client time not as expected")
	assert.Equal(t, 1, a.SubKey.KeyType, "Subkey type not as expected")
	assert.Equal(t, []byte("12345678"), a.SubKey.KeyValue, "Subkey value not as expected")
	assert.Equal(t, 2, len(a.AuthorizationData), "Number of Authorization data items not as expected")
	for i, entry := range a.AuthorizationData {
		assert.Equal(t, 1, entry.ADType, fmt.Sprintf("Authorization type of entry %d not as expected", i+1))
		assert.Equal(t, []byte("foobar"), entry.ADData, fmt.Sprintf("Authorization data of entry %d not as expected", i+1))
	}
}

func TestUnmarshalAuthenticator_optionalsempty(t *testing.T) {
	a := unmarshalAuthenticator_test(t, "encode_krb5_authenticator(optionalsempty)")
	//Parse the test time value into a time.Time type
	tt, _ := time.Parse(testdata.TEST_TIME_FORMAT, testdata.TEST_TIME)

	assert.Equal(t, testdata.TEST_KVNO, a.AVNO, "Authenticator version number not as expected")
	assert.Equal(t, testdata.TEST_REALM, a.CRealm, "CRealm not as expected")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMETYPE, a.CName.NameType, "CName NameType not as expected")
	assert.Equal(t, len(testdata.TEST_PRINCIPALNAME_NAMESTRING), len(a.CName.NameString), "CName does not have the expected number of NameStrings")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMESTRING, a.CName.NameString, "CName entries not as expected")
	assert.Equal(t, 123456, a.Cusec, "Client microseconds not as expected")
	assert.Equal(t, tt, a.CTime, "Client time not as expected")
}

func TestUnmarshalAuthenticator_optionalsNULL(t *testing.T) {
	a := unmarshalAuthenticator_test(t, "encode_krb5_authenticator(optionalsNULL)")
	//Parse the test time value into a time.Time type
	tt, _ := time.Parse(testdata.TEST_TIME_FORMAT, testdata.TEST_TIME)

	assert.Equal(t, testdata.TEST_KVNO, a.AVNO, "Authenticator version number not as expected")
	assert.Equal(t, testdata.TEST_REALM, a.CRealm, "CRealm not as expected")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMETYPE, a.CName.NameType, "CName NameType not as expected")
	assert.Equal(t, len(testdata.TEST_PRINCIPALNAME_NAMESTRING), len(a.CName.NameString), "CName does not have the expected number of NameStrings")
	assert.Equal(t, testdata.TEST_PRINCIPALNAME_NAMESTRING, a.CName.NameString, "CName entries not as expected")
	assert.Equal(t, 123456, a.Cusec, "Client microseconds not as expected")
	assert.Equal(t, tt, a.CTime, "Client time not as expected")
}
