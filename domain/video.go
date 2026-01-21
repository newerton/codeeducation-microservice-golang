package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Video godoc
type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

// NewVideo godoc
func NewVideo() *Video {
	return &Video{}
}

// Validate godoc
func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}

// // Unmarshal godoc
// func (video *Video) Unmarshal(payload []byte) Video {
// 	err := json.Unmarshal(payload, video)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return *video
// }

// // Download godoc
// func (video *Video) Download(bucketName string, storagePath string) (Video, error) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 		return *video, err
// 	}

// 	bkt := client.Bucket(bucketName)
// 	obj := bkt.Object(video.Path)
// 	r, err := obj.NewReader(ctx)
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 		return *video, err
// 	}

// 	defer r.Close()

// 	body, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 		return *video, err
// 	}

// 	f, err := os.Create(storagePath + "/" + video.UUID + ".mp4")
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 		return *video, err
// 	}

// 	_, err = f.Write(body)
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 		return *video, err
// 	}

// 	defer f.Close()

// 	fmt.Println("Video", video.UUID, "has been Stored")
// 	return *video, nil

// }

// // Fragment godoc
// func (video *Video) Fragment(storagePath string) Video {
// 	err := os.Mkdir(storagePath+"/"+video.UUID, os.ModePerm)
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println("Gerando fragmento:", video.UUID)
// 	source := storagePath + "/" + video.UUID + ".mp4"
// 	target := storagePath + "/" + video.UUID + ".frag"

// 	cmd := exec.Command("mp4fragment", source, target)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 	}

// 	printOutput(output)

// 	return *video
// }

// // Encode godoc
// func (video *Video) Encode(storagePath string) Video {

// 	cmdArgs := []string{}
// 	cmdArgs = append(cmdArgs, storagePath+"/"+video.UUID+".frag")
// 	cmdArgs = append(cmdArgs, "--use-segment-timeline")
// 	cmdArgs = append(cmdArgs, "--o")
// 	cmdArgs = append(cmdArgs, storagePath+"/"+video.UUID)
// 	cmdArgs = append(cmdArgs, "--f")
// 	cmdArgs = append(cmdArgs, "--exec-dir")
// 	cmdArgs = append(cmdArgs, "/usr/local/bin")

// 	cmd := exec.Command("mp4dash", cmdArgs...)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		video.Status = "error"
// 		fmt.Println(err.Error())
// 	}

// 	printOutput(output)

// 	return *video
// }

// // UploadObject godoc
// func (video *Video) UploadObject(completePath string, storagePath string, bucketName string, client *storage.Client, ctx context.Context) error {
// 	path := strings.Split(completePath, storagePath+"/")

// 	f, err := os.Open(completePath)
// 	if err != nil {
// 		fmt.Println("Error during the upload", err.Error())
// 		return err
// 	}
// 	defer f.Close()

// 	wc := client.Bucket(bucketName).Object(path[1]).NewWriter(ctx)
// 	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

// 	if _, err = io.Copy(wc, f); err != nil {
// 		return err
// 	}

// 	if err := wc.Close(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Finish godoc
// func (video *Video) Finish(storagePath string) {
// 	err := os.Remove(storagePath + "/" + video.UUID + ".mp4")
// 	if err != nil {
// 		fmt.Printf("Error removing MP4: ", video.UUID, ".mp4")
// 	}

// 	err = os.Remove(storagePath + "/" + video.UUID + ".frag")
// 	if err != nil {
// 		fmt.Printf("Error removing FRAG: ", video.UUID, ".frag")
// 	}

// 	err = os.RemoveAll(storagePath + "/" + video.UUID)
// 	if err != nil {
// 		fmt.Printf("Error removing FOLDER: ", video.Path)
// 	}

// 	fmt.Println("Files has been removed", video.UUID)
// }

// // GetVideoPaths godoc
// func (video *Video) GetVideoPaths() []string {
// 	var paths []string
// 	filepath.Walk("/tmp/abc-123-def-456", func(path string, info os.FileInfo, err error) error {
// 		paths = append(paths, path)
// 		return nil
// 	})
// 	return paths
// }

// // printOutput godoc
// func printOutput(out []byte) {
// 	if len(out) > 0 {
// 		fmt.Printf("===> Output: %s\n", string(out))
// 	}
// }
