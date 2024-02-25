package roadmap_ps_repositories

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/un-defined-gsc/un-defined-backend/internal/config"
	base_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/base"
	roadmap_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/roadmap"
	"github.com/un-defined-gsc/un-defined-backend/pkg/db_adapters"
)

var dbpool *pgxpool.Pool
var roadmapRepo *roadmapRepository

func TestMain(t *testing.T) {
	err := config.InitializeConfig("../../../../config/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.GetConfig()
	cfg.Database.Host = "localhost"
	dbpool, err = db_adapters.NewPostgressClient(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Database)
	if err != nil {
		t.Fatal(err)
	}
	roadmapRepo = &roadmapRepository{
		dbpool: dbpool,
	}
	testCRUD_Roadmap(t)
}

func testCRUD_Roadmap(t *testing.T) {
	mockRoadmapObj := &roadmap_domain.Roadmap{
		Name:        "Test Roadmap",
		Description: "Test Description",
		Base: base_domain.Base{
			ID: uuid.New(),
		},
	}
	testCreate := t.Run("CRUD Create",
		func(t *testing.T) {
			// Test for Create method
			t.Logf("Creating Roadmap: %v", mockRoadmapObj)
			err := roadmapRepo.Create(context.Background(), mockRoadmapObj)
			if err != nil {
				t.Error(err)
			}
		})
	if !testCreate {
		t.Error("TestCreate failed")
	}
	testUpdate := t.Run("CRUD Update",
		func(t *testing.T) {
			// Test for Update method
			mockRoadmapObj.Name = "Test Roadmap Updated"
			mockRoadmapObj.Description = "Test Description Updated"

			t.Logf("Updating Roadmap: %v", mockRoadmapObj)
			err := roadmapRepo.Update(context.Background(), mockRoadmapObj)
			if err != nil {
				t.Error(err)
			}
		})
	if !testUpdate {
		t.Error("TestUpdate failed")
	}
	testFilter := t.Run("CRUD Filter",
		func(t *testing.T) {
			// Test for Filter method
			t.Logf("Filtering Roadmap: %v", mockRoadmapObj)
			_, err := roadmapRepo.Filter(context.Background(), mockRoadmapObj)
			if err != nil {
				t.Error(err)
			}
		})
	if !testFilter {
		t.Error("TestFilter failed")
	}
	testDelete := t.Run("CRUD Delete",
		func(t *testing.T) {
			// Test for Delete method
			t.Logf("Deleting Roadmap: %v", mockRoadmapObj)
			err := roadmapRepo.Delete(context.Background(), mockRoadmapObj.ID)
			if err != nil {
				t.Error(err)
			}
		},
	)
	if !testDelete {
		t.Error("TestDelete failed")
	}

}
