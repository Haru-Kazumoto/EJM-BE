package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type BeError struct {
	status    int
	id        string
	localizer echo.Context
	msg       *i18n.Message
}

func (e *BeError) SetLocalizer(c echo.Context) {
	e.localizer = c
}

func (e *BeError) Error() string {
	translate, _ := Translate(e.localizer, e.msg)

	return translate
}

func (e *BeError) GetIDTranslate() string {
	return e.id
}

func (e *BeError) GetCode() int {
	return e.status
}

var (
	ErrUserNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "user.not.found",
			Other: "User Not Found",
		},
	}

	ErrMappingCodeNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "user.not.found",
			Other: "User Not Found",
		},
	}

	ErrMappingKeywordListNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "mappingekywordlist.not.found",
			Other: "Mapping Keyword Lisr Not Found",
		},
	}

	ErrListOpCodeNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "liatopcode.not.found",
			Other: "List Op Code Not Found",
		},
	}

	ErrJenisTransaksiNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "JenisTransaksi.not.found",
			Other: "JenisTransaksi Not Found",
		},
	}
	ErrLookupNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "lookup.not.found",
			Other: "Data Not Found",
		},
	}

	ErrUsernameExist = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "username.already.exist",
			Other: "Username already exist",
		},
	}
	ErrRoleNotExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "role.not.found",
			Other: "Role Not Found",
		},
	}
	
	ErrRoleAlreadyExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "role.already.exists",
			Other: "Role already exists",
		},
	}
	ErrMappingCodeNotExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "mappingCode.not.found",
			Other: "Mapping Code Not Found",
		},
	}
	ErrDefinitionAlreadyExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "mappingCode.already.exists",
			Other: "Mapping Code already exists",
		},
	}
	ErrMappingKeywordListAlreadyExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "mappingKeywroddList.already.exists",
			Other: "Mapping Keyword List already exists",
		},
	}
	ErrTransactionTypeAlreadyExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "TransactionType.already.exists",
			Other: "Transaction type already exists",
		},
	}
	ErrListOpCodeAlreadyExists = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "listOpCode.already.exists",
			Other: "List Op Code already exists",
		},
	}
	ErrCredentialInvalid = &BeError{
		status: http.StatusForbidden,
		msg: &i18n.Message{
			ID:    "credential.invalid",
			Other: "User Credential Invalid",
		},
	}

	ErrForbiddenLoginByLink = &BeError{
		status: http.StatusForbidden,
		msg: &i18n.Message{
			ID:    "forbidden.login.bylink",
			Other: "This user cannot login by link, contact administrator to open the access",
		},
	}
	ErrUnauthorizedRole = &BeError{
		status: http.StatusForbidden,
		msg: &i18n.Message{
			ID:    "role.forbidden.access",
			Other: "Role doesn't have access",
		},
	}
	ErrTokenInvalid = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "invalid.token",
			Other: "Invalid Token",
		},
	}

	ErrUserUnauhotirzed = &BeError{
		status: http.StatusUnauthorized,
		msg: &i18n.Message{
			ID:    "user.unauthorized",
			Other: "User Unauthenticated",
		},
	}

	ErrFileTooLarge = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "file.too.large",
			Other: "File too large, max 2mb",
		},
	}
	ErrFileHasNoExtension = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "file.has.no.extension",
			Other: "File has not extension",
		},
	}
	ErrFileExtension = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "file.extension.forbidden",
			Other: "File format invalid",
		},
	}

	ErrDataVolumeNotFound = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "volume.data.notfound",
			Other: "Volume Data Not Found",
		},
	}

	ErrRoleCannotDeleted = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "roles.default",
			Other: "Role Default Cannot Be Deleted",
		},
	}
	ErrCompanyAlreadyExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "company.exists",
			Other: "Company already exists",
		},
	}
	ErrCompanyStructuresEmpty = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "company_structures.empty",
			Other: "Company Structures must be set",
		},
	}
	ErrCompanyNotExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "company.not.exists",
			Other: "Company not exists",
		},
	}

	ErrQuestionTemplateAlreadyExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "question_templates.exists",
			Other: "Question Template already exists",
		},
	}

	ErrQuestionTemplateNotExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "question_templates.not.exists",
			Other: "Question Template Not exists",
		},
	}

	ErrQuestionEmpty = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "question.empty",
			Other: "Question Empty",
		},
	}

	ErrQuestionNotExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "question.not_exists",
			Other: "Question Not Exists",
		},
	}

	ErrQuestionAlreadyExists = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "question.exists",
			Other: "Question Exists",
		},
	}

	ErrProjectNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "projects.not_found",
			Other: "Project Not Found",
		},
	}

	ErrInventoryNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "inventory.not_found",
			Other: "Inventory Not Found",
		},
	}

	ErrCompanyStructureNotFound = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "structure.not.found",
			Other: "Structure Not Found",
		},
	}

	ErrAllUsersAlreadyAssigned = &BeError{
		status: http.StatusNotFound,
		msg: &i18n.Message{
			ID:    "users.all.assigned",
			Other: "All Users Already Assigned",
		},
	}
	ErrMenuCannotLessThanOne = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "menu.empty",
			Other: "Menu Cannot less than one",
		},
	}
	ErrCannotDispatchJobQueue = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "jobqueue.error",
			Other: "Cannot dispatch job queue",
		},
	}
	ErrCreateProjectLog = &BeError{
		status: http.StatusBadRequest,
		msg: &i18n.Message{
			ID:    "create.project.log.error",
			Other: "Error Create Project Log",
		},
	}
)
