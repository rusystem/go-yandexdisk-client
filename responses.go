package yandexdisk

import (
	"fmt"
	"time"
)

type ErrorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

type SuccessResponse struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

type Disk struct {
	MaxFileSize                int           `json:"max_file_size,omitempty"`
	PaidMaxFileSize            int64         `json:"paid_max_file_size,omitempty"`
	TotalSpace                 int64         `json:"total_space,omitempty"`
	TrashSize                  int           `json:"trash_size,omitempty"`
	IsPaid                     bool          `json:"is_paid,omitempty"`
	UsedSpace                  int64         `json:"used_space,omitempty"`
	SystemFolders              SystemFolders `json:"system_folders,omitempty"`
	User                       User          `json:"user,omitempty"`
	UnlimitedAutouploadEnabled bool          `json:"unlimited_autoupload_enabled,omitempty"`
	Revision                   int64         `json:"revision,omitempty"`
}

type SystemFolders struct {
	Odnoklassniki string `json:"odnoklassniki,omitempty"`
	Google        string `json:"google,omitempty"`
	Instagram     string `json:"instagram,omitempty"`
	Vkontakte     string `json:"vkontakte,omitempty"`
	Mailru        string `json:"mailru,omitempty"`
	Downloads     string `json:"downloads,omitempty"`
	Applications  string `json:"applications,omitempty"`
	Facebook      string `json:"facebook,omitempty"`
	Social        string `json:"social,omitempty"`
	Scans         string `json:"scans,omitempty"`
	Screenshots   string `json:"screenshots,omitempty"`
	Photostream   string `json:"photostream,omitempty"`
}

type User struct {
	Country     string `json:"country,omitempty"`
	Login       string `json:"login,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	UID         string `json:"uid,omitempty"`
}

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

func (e *ErrorResponse) Info() string {
	return fmt.Sprintf("message: %s, description: %s, error: %s\n", e.Message, e.Description, e.Error)
}

func (d *Disk) Info() string {
	return fmt.Sprintf("User name: %s\nUsed memory: %d of %d bytes\n", d.User.DisplayName, d.UsedSpace, d.TotalSpace)
}

func (i *Item) Info() string {
	return fmt.Sprintf("filename: %s | filesize: %d | filepath: %s\n", i.Name, i.Size, i.Path)
}
