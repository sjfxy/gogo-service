package service

import (
	"testing"

	"github.com/cloudnativego/cfmgo"
	"github.com/cloudnativego/gogo-engine"
	"github.com/sjfxy/gogo-service/fakes"
)

var (
	fakeDBURI = "mongodb://sin:sj930826@ds149875.mlab.com:49875/CloudFoundry_e4r3t4m5_n6hgr8tt"
)

func TestAddMatchShowsUpInMongoRepository(t *testing.T) {
	var fakeMatches = []matchRecord{}
	var matchesCollection = cfmgo.Connect(
		fakes.FakeNewCollectionDialer(fakeMatches),
		fakeDBURI,
		MatchesCollectionName)

	repo := newMongoMatchRepository(matchesCollection)
	match := gogo.NewMatch(19, "bob", "alfred")
	err := repo.addMatch(match)
	if err != nil {
		t.Errorf("Error adding match to mongo: %v", err)
	}

	matches, err := repo.getMatches()
	if err != nil {
		t.Errorf("Got an error retrieving matches: %v", err)
	}
	if len(matches) != 1 {
		t.Errorf("Expected matches length to be 1; received %d", len(matches))
	}
}

func TestGetMatchRetrievesProperMatchFromMongo(t *testing.T) {
	fakes.TargetCount = 1
	var fakeMatches = []matchRecord{}
	var matchesCollection = cfmgo.Connect(
		fakes.FakeNewCollectionDialer(fakeMatches),
		fakeDBURI,
		MatchesCollectionName)

	repo := newMongoMatchRepository(matchesCollection)
	match := gogo.NewMatch(19, "bob", "alfred")
	err := repo.addMatch(match)
	if err != nil {
		t.Errorf("Error adding match to mongo: %v", err)
	}

	targetID := match.ID
	foundMatch, err := repo.getMatch(targetID)
	if err != nil {
		t.Errorf("Unable to find match with ID: %v... %s", targetID, err)
	}

	if foundMatch.GridSize != 19 || foundMatch.PlayerBlack != "bob" {
		t.Errorf("Unexpected match results: %v", foundMatch)
	}
}

func TestGetNonExistentMatchReturnsError(t *testing.T) {
	fakes.TargetCount = 0
	var fakeMatches = []matchRecord{}
	var matchesCollection = cfmgo.Connect(
		fakes.FakeNewCollectionDialer(fakeMatches),
		fakeDBURI,
		MatchesCollectionName)

	repo := newMongoMatchRepository(matchesCollection)

	_, err := repo.getMatch("bad_id")
	if err == nil {
		t.Errorf("Expected getMatch to error with incorrect match details")
	}

	if err.Error() != "Match not found" {
		t.Errorf("Expected 'Match not found' error; received: '%v'", err)
	}
}
