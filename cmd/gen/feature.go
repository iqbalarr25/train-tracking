package gen_cmd

import (
	"TrainTracking/pkg/logger"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var featureName string

var GenerateFeatureCMD = &cobra.Command{
	Use:   "feature",
	Short: "Generate feature scaffolding",
	Long: `Usage:
	feature <name>`,
	Run: generateFeatureScaffold,
}

func init() {
}

func generateFeatureScaffold(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		featureName = args[0]
	}
	if featureName == "" {
		logger.Panic("Failed to generate feature, plesase provide a name")
	}

	generateFeatInterface(featureName)
	generateFeatRequest(featureName)
	generateFeatResponse(featureName)
	generateFeatService(featureName)
	generateFeatHandler(featureName)

	logger.Info("Successfully generated %s feature. \nHappy coding! \nDon't forget to take a break for some caffeine 🚀", featureName)
}

func generateFeatInterface(name string) {
	lowerName := strings.ToLower(name)
	fileDir := "internal/features/" + lowerName
	filename := fmt.Sprintf("%s.interface.go", lowerName)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create feature interface directory: %v", err)
		}
	}

	filePath := filepath.Join(fileDir, filename)
	fileContent := fmt.Sprintf(`package %s

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	Index(ctx *fiber.Ctx) error
}

type IService interface {
	Fetch(request *%[2]sRequest) (*%[2]sResponse, error)
}`, lowerName, name)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create interface file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ %s feature interface successfully created at %s.", name, filePath))
}

func generateFeatHandler(name string) {
	lowerName := strings.ToLower(name)
	fileDir := "internal/features/" + lowerName
	filename := fmt.Sprintf("%s.handler.go", lowerName)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create handler directory: %v", err)
		}
	}

	filePath := filepath.Join(fileDir, filename)
	fileContent := fmt.Sprintf(`package %s

type %[2]sHandler struct {
	service IService
	cfg     config.App
}

func New%[2]sHandler() IHandler {
	return &%[2]sHandler{
		New%[2]sService(),
		config.GetApp(),
	}
}

func (h *%[2]sHandler) Index(ctx *fiber.Ctx) error {
	results, err := h.service.Fetch(&%[2]sRequest{})
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.Response{
			Code:    fiber.StatusOK,
			Success: true,
			Message: "Failed",
			Data:    results,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.Response{
		Code:    fiber.StatusOK,
		Success: true,
		Message: "Success",
		Data:    results,
	})
}`, lowerName, name)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create migration file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ %s feature handler successfully created at %s.", name, filePath))
}

func generateFeatRequest(name string) {
	lowerName := strings.ToLower(name)
	fileDir := "internal/features/" + lowerName
	filename := fmt.Sprintf("%s.request.go", lowerName)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create feature interface directory: %v", err)
		}
	}

	filePath := filepath.Join(fileDir, filename)
	fileContent := fmt.Sprintf(`package %s

type %[2]sRequest struct {
	// Your code
}

// func ValidateStore(jsonString string) (*%[2]sRequest, error) {
// 	validator := validator.New()
// 	request := &%[2]sRequest{}
// 	if err := json.Unmarshal([]byte(jsonString), &request); err != nil {
// 		helper.Exception(err)
// 		return nil, errors.New("failed decode payload")
// 	}

// 	if err := validator.Struct(request); err != nil {
// 		helper.Exception(err)
// 		return nil, err
// 	}

// 	return request, nil
// }
`, lowerName, name)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create request file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ %s feature request successfully created at %s.", name, filePath))
}

func generateFeatService(name string) {
	lowerName := strings.ToLower(name)
	fileDir := "internal/features/" + lowerName
	filename := fmt.Sprintf("%s.service.go", lowerName)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create feature response directory: %v", err)
		}
	}

	filePath := filepath.Join(fileDir, filename)
	fileContent := fmt.Sprintf(`package %s

type %[2]sService struct {
	db  *gorm.DB
	cfg config.App
}

func (service *%[2]sService) Fetch(request *%[2]sRequest) (*%[2]sResponse, error) {
	return &%[2]sResponse{}, nil
}

func New%[2]sService() IService {
	return &%[2]sService{
		config.GetDBConnection(),
		config.GetApp(),
	}
}`, lowerName, name)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create service file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ %s feature service successfully created at %s.", name, filePath))
}

func generateFeatResponse(name string) {
	lowerName := strings.ToLower(name)
	fileDir := "internal/features/" + lowerName
	filename := fmt.Sprintf("%s.response.go", lowerName)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create feature response directory: %v", err)
		}
	}

	filePath := filepath.Join(fileDir, filename)
	fileContent := fmt.Sprintf(`package %s

type %[2]sResponse struct {
	// your response
}`, lowerName, name)

	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		logger.Fatal("Failed to create response file: %v", err)
	}

	logger.Info(fmt.Sprintf("✅ %s feature response successfully created at %s.", name, filePath))
}
