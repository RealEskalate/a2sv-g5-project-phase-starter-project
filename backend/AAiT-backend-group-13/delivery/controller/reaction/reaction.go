package reactioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/group13/blog/delivery/common"
	basecontroller "github.com/group13/blog/delivery/controller/base"
	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	reactioncmd "github.com/group13/blog/usecase/reaction/command"
)

type ReactionController struct {
	basecontroller.BaseHandler
	updateReactionHandler icmd.IHandler[*reactioncmd.UpdateCommand, bool]
	deleteReactionHandler icmd.IHandler[*reactioncmd.DeleteCommand, bool]
}

var _ common.IController = &ReactionController{}

type Config struct {
	UpdateReactionHandler icmd.IHandler[*reactioncmd.UpdateCommand, bool]
	DeleteReactionHandler icmd.IHandler[*reactioncmd.DeleteCommand, bool]
}

func New(config Config) *ReactionController {
	return &ReactionController{
		updateReactionHandler: config.UpdateReactionHandler,
		deleteReactionHandler: config.DeleteReactionHandler,
	}
}

// RegisterPublic registers public routes.
func (r *ReactionController) RegisterPublic(route *gin.RouterGroup) {

}

// RegisterPrivileged registers privileged routes.
func (r *ReactionController) RegisterPrivileged(route *gin.RouterGroup) {}

// RegisterProtected registers protected routes.
func (r *ReactionController) RegisterProtected(route *gin.RouterGroup) {
	reaction := route.Group("/blogs")
	{

		reaction.PUT("/:id/reaction", r.UpdateReaction)
		reaction.DELETE("/:id/reaction", r.DeleteReaction)

	}

}

func (r ReactionController) UpdateReaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "blogId is not valid"})
		return
	}
	var reaction ReactionDto

	if err := c.ShouldBindJSON(&reaction); err != nil {
		r.Respond(c, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	decoded, err := uuid.Parse(reaction.UserId)

	if err != nil {
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "userId is not valid"})
		return
	}

	command := reactioncmd.NewUpdateCommand(reaction.IsLike, id, uuid.UUID(decoded))

	_, err = r.updateReactionHandler.Handle(command)

	if err != nil {
		r.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	r.Respond(c, http.StatusNoContent, gin.H{})

}

func (r ReactionController) DeleteReaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "blogId is not valid"})
		return
	}
	var reaction ReactionDto

	if err := c.ShouldBindJSON(&reaction); err != nil {
		r.Respond(c, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	decoded, err := uuid.Parse(reaction.UserId)

	if err != nil {
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "userId is not valid"})
		return
	}

	command := reactioncmd.NewDeleteCommand(id, uuid.UUID(decoded))

	_, err = r.deleteReactionHandler.Handle(command)

	if err != nil {
		r.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	r.Respond(c, http.StatusNoContent, gin.H{})
}
