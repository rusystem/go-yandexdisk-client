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
	MaxFileSize                int           `json:"max_file_size"`
	PaidMaxFileSize            int64         `json:"paid_max_file_size"`
	TotalSpace                 int64         `json:"total_space"`
	TrashSize                  int           `json:"trash_size"`
	IsPaid                     bool          `json:"is_paid"`
	UsedSpace                  int64         `json:"used_space"`
	SystemFolders              SystemFolders `json:"system_folders"`
	User                       User          `json:"user"`
	UnlimitedAutouploadEnabled bool          `json:"unlimited_autoupload_enabled"`
	Revision                   int64         `json:"revision"`
}

type SystemFolders struct {
	Odnoklassniki string `json:"odnoklassniki"`
	Google        string `json:"google"`
	Instagram     string `json:"instagram"`
	Vkontakte     string `json:"vkontakte"`
	Mailru        string `json:"mailru"`
	Downloads     string `json:"downloads"`
	Applications  string `json:"applications"`
	Facebook      string `json:"facebook"`
	Social        string `json:"social"`
	Scans         string `json:"scans"`
	Screenshots   string `json:"screenshots"`
	Photostream   string `json:"photostream"`
}

type User struct {
	Country     string `json:"country"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	UID         string `json:"uid"`
}

type FilesResourceList struct {
	Items  []Item `json:"items"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type Item struct {
	AntivirusStatus string     `json:"antivirus_status"`
	Size            int64      `json:"size"`
	CommentIds      CommentIds `json:"comment_ids"`
	Name            string     `json:"name"`
	Exif            Exif       `json:"exif"`
	Created         time.Time  `json:"created"`
	ResourceID      string     `json:"resource_id"`
	Modified        time.Time  `json:"modified"`
	MimeType        string     `json:"mime_type"`
	File            string     `json:"file"`
	Path            string     `json:"path"`
	MediaType       string     `json:"media_type"`
	Sha256          string     `json:"sha256"`
	Type            string     `json:"type"`
	Md5             string     `json:"md5"`
	Revision        int64      `json:"revision"`
}

type CommentIds struct {
	PrivateResource string `json:"private_resource"`
	PublicResource  string `json:"public_resource"`
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
