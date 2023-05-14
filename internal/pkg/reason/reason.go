package reason

/*
By https://github.com/answerdev/answer
*/

const (
	// Success .
	Success = "base.success"
	// UnknownError unknown error
	UnknownError = "base.unknown"
	// RequestFormatError request format error
	RequestFormatError = "base.request_format_error"
	// UnauthorizedError unauthorized error
	UnauthorizedError = "base.unauthorized_error"
	// DatabaseError database error
	DatabaseError = "base.database_error"
	// ServerError server error
	ServerError = "base.server_error"
)

const (
	FollowAlreadyExist = "error.follow.already_exist"
	FollowNotFound     = "error.follow.not_found"

	AuthHeaderInvalid                = "error.auth.header_invalid"
	TokenInvalid                     = "error.token.invalid"
	TokenNotExistOrExpired           = "error.token.not_exist_or_expired"
	EmailOrPasswordWrong             = "error.object.email_or_password_incorrect"
	CommentNotFound                  = "error.comment.not_found"
	CommentCannotEditAfterDeadline   = "error.comment.cannot_edit_after_deadline"
	QuestionNotFound                 = "error.question.not_found"
	QuestionCannotDeleted            = "error.question.cannot_deleted"
	QuestionCannotClose              = "error.question.cannot_close"
	QuestionCannotUpdate             = "error.question.cannot_update"
	QuestionAlreadyDeleted           = "error.question.already_deleted"
	AnswerNotFound                   = "error.answer.not_found"
	AnswerCannotDeleted              = "error.answer.cannot_deleted"
	AnswerCannotUpdate               = "error.answer.cannot_update"
	AnswerCannotAddByClosedQuestion  = "error.answer.question_closed_cannot_add"
	CommentEditWithoutPermission     = "error.comment.edit_without_permission"
	DisallowVote                     = "error.object.disallow_vote"
	DisallowFollow                   = "error.object.disallow_follow"
	DisallowVoteYourSelf             = "error.object.disallow_vote_your_self"
	CaptchaVerificationFailed        = "error.object.captcha_verification_failed"
	OldPasswordVerificationFailed    = "error.object.old_password_verification_failed"
	NewPasswordSameAsPreviousSetting = "error.object.new_password_same_as_previous_setting"
	UserNotFound                     = "error.user.not_found"
	UsernameInvalid                  = "error.user.username_invalid"
	UsernameDuplicate                = "error.user.username_duplicate"
	UserSetAvatar                    = "error.user.set_avatar"
	EmailDuplicate                   = "error.email.duplicate"
	EmailShouldRequestLater          = "error.email.should_request_later"
	EmailCodeExpired                 = "error.email.code_expired"
	EmailCodeIncorrect               = "error.email.code_incorrect"
	EmailVerifyURLExpired            = "error.email.verify_url_expired"
	EmailNeedToBeVerified            = "error.email.need_to_be_verified"
	EmailNeedToBeVerifiedBeforeAct   = "error.email.need_to_be_verified_before_action"
	UserSuspended                    = "error.user.suspended"
	ObjectNotFound                   = "error.object.not_found"
	TagNotFound                      = "error.tag.not_found"
	TagNotContainSynonym             = "error.tag.not_contain_synonym_tags"
	TagCannotUpdate                  = "error.tag.cannot_update"
	TagIsUsedCannotDelete            = "error.tag.is_used_cannot_delete"
	TagAlreadyExist                  = "error.tag.already_exist"
	TagDuplicateInObeject            = "error.tag.duplicate_in_object"
	PostNotFound                     = "error.post.not_found"
	RankFailToMeetTheCondition       = "error.rank.fail_to_meet_the_condition"
	LikeAlreadyExist                 = "error.like.already_exist"
	LikeCancelFailToNotLiked         = "error.like.cancel_fail_to_not_liked"
	VoteRankFailToMeetTheCondition   = "error.rank.vote_fail_to_meet_the_condition"
	ThemeNotFound                    = "error.theme.not_found"
	LangNotFound                     = "error.lang.not_found"
	ReportHandleFailed               = "error.report.handle_failed"
	ReportNotFound                   = "error.report.not_found"
	ReadConfigFailed                 = "error.config.read_config_failed"
	DatabaseConnectionFailed         = "error.database.connection_failed"
	InstallCreateTableFailed         = "error.database.create_table_failed"
	InstallConfigFailed              = "error.install.create_config_failed"
	SiteInfoNotFound                 = "error.site_info.not_found"
	UploadFileSourceUnsupported      = "error.upload.source_unsupported"
	UploadFileUnsupportedFileFormat  = "error.upload.unsupported_file_format"
	RecommendTagNotExist             = "error.tag.recommend_tag_not_found"
	RecommendTagEnter                = "error.tag.recommend_tag_enter"
	RevisionReviewUnderway           = "error.revision.review_underway"
	RevisionNoPermission             = "error.revision.no_permission"
	UserCannotUpdateYourRole         = "error.user.cannot_update_your_role"
	TagCannotSetSynonymAsItself      = "error.tag.cannot_set_synonym_as_itself"
	NotAllowedRegistration           = "error.user.not_allowed_registration"
	SMTPConfigFromNameCannotBeEmail  = "error.smtp.config_from_name_cannot_be_email"
	AdminCannotUpdateTheirPassword   = "error.admin.cannot_update_their_password"
	AdminCannotModifySelfStatus      = "error.admin.cannot_modify_self_status"
)
