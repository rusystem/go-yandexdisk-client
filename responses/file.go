package responses

import (
	"fmt"
	"time"
)

type FilesResourceList struct {
	Items  []Item `json:"items,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type Item struct {
	AntivirusStatus string     `json:"antivirus_status,omitempty"`
	Size            int64      `json:"size,omitempty"`
	CommentIds      CommentIds `json:"comment_ids,omitempty"`
	Name            string     `json:"name,omitempty"`
	Exif            Exif       `json:"exif,omitempty"`
	Created         time.Time  `json:"created,omitempty"`
	ResourceID      string     `json:"resource_id,omitempty"`
	Modified        time.Time  `json:"modified,omitempty"`
	MimeType        string     `json:"mime_type,omitempty"`
	File            string     `json:"file,omitempty"`
	Path            string     `json:"path,omitempty"`
	MediaType       string     `json:"media_type,omitempty"`
	Sha256          string     `json:"sha256,omitempty"`
	Type            string     `json:"type,omitempty"`
	Md5             string     `json:"md5,omitempty"`
	Revision        int64      `json:"revision,omitempty"`
}

type CommentIds struct {
	PrivateResource string `json:"private_resource,omitempty"`
	PublicResource  string `json:"public_resource,omitempty"`
}

type Exif struct {
	DateTime string `json:"date_time"`
}

func (i *Item) Info() string {
	return fmt.Sprintf("filename: %s | filesize: %d | filepath: %s\n", i.Name, i.Size, i.Path)
}
