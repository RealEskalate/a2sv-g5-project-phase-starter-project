package reactioncontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	basecontroller "github.com/group13/blog/delivery/base"
	common "github.com/group13/blog/delivery/common/icontroller"
	"github.com/group13/blog/delivery/controller/dto"
	er "github.com/group13/blog/domain/errors"
	icmd "github.com/group13/blog/usecase/common/cqrs/command"
	updatereaction "github.com/group13/blog/usecase/reaction/command/update"
)

type ReactionController struct {
	BaseController        basecontroller.BaseHandler
	updateReactionHandler icmd.IHandler[*updatereaction.Command, bool]
	deleteReactionHandler icmd.IHandler[uuid.UUID, bool]
}

var _ common.IController = &ReactionController{}

type Config struct {
	updateReactionHandler icmd.IHandler[*updatereaction.Command, bool]
	deleteReactionHandler icmd.IHandler[uuid.UUID, bool]
	basecontroller        basecontroller.BaseHandler
}

func New(config Config) *ReactionController {
	return &ReactionController{
		updateReactionHandler: config.updateReactionHandler,
		deleteReactionHandler: config.deleteReactionHandler,
		BaseController:        config.basecontroller,
	}
}

// RegisterPublic registers public routes.
func (r *ReactionController) RegisterPublic(route *gin.RouterGroup) {}

// RegisterPrivileged registers privileged routes.
func (r *ReactionController) RegisterPrivileged(route *gin.RouterGroup) {}

// RegisterProtected registers protected routes.
func (r *ReactionController) RegisterProtected(route *gin.RouterGroup) {
	reaction := route.Group("/reaction")
	{

		reaction.PUT("/:id", r.UpdateReaction)
		reaction.DELETE("/:id", r.DeleteReaction)

	}
}

func (r ReactionController) UpdateReaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		r.BaseController.Respond(c, http.StatusBadRequest, gin.H{"error": "Id is not valid"})
		return
	}
	var reaction dto.ReactionDto

	if err := c.ShouldBindJSON(&reaction); err != nil {
		r.BaseController.Respond(c, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	command := updatereaction.NewCommand(reaction.IsLike, id, reaction.UserId)

	_, err = r.updateReactionHandler.Handle(command)

	if err != nil {
		r.BaseController.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	r.BaseController.Respond(c, http.StatusNoContent, gin.H{})

}

func (r ReactionController) DeleteReaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		r.BaseController.Respond(c, http.StatusBadRequest, gin.H{"error": "Id is not valid"})
		return
	}

	_, err = r.deleteReactionHandler.Handle(id)

	if err != nil {
		r.BaseController.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	r.BaseController.Respond(c, http.StatusNoContent, gin.H{})

}
