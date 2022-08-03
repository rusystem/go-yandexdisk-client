package responses

import "fmt"

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

func (d *Disk) Info() string {
	return fmt.Sprintf("User name: %s\nUsed memory: %d of %d bytes\n", d.User.DisplayName, d.UsedSpace, d.TotalSpace)
}
