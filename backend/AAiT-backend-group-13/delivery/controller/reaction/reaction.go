package reactioncontroller

import (
	
	"log"
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
	UpdateReactionHandler icmd.IHandler[*reactioncmd.UpdateCommand, bool]
	DeleteReactionHandler icmd.IHandler[uuid.UUID, bool]
}

var _ common.IController = &ReactionController{}

type Config struct {
	UpdateReactionHandler icmd.IHandler[*reactioncmd.UpdateCommand, bool]
	DeleteReactionHandler icmd.IHandler[uuid.UUID, bool]
}

func New(config Config) *ReactionController {
	return &ReactionController{
		UpdateReactionHandler: config.UpdateReactionHandler,
		DeleteReactionHandler: config.DeleteReactionHandler,
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
	log.Println("UpdateReaction started")

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("Invalid UUID:", err)
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "Id is not valid"})
		return
	}
	
	// log.Printf("Binding JSON %s", )
	var reaction ReactionDto
	if err := c.ShouldBindJSON(&reaction); err != nil {
		log.Println("Failed to bind JSON:", err)
		r.Respond(c, http.StatusBadRequest, er.NewBadRequest(err.Error()))
		return
	}

	log.Println("Creating update command with is_like:", reaction.IsLike, "user_id:", reaction.UserId)
	command := reactioncmd.NewUpdateCommand(reaction.IsLike, id, reaction.UserId)

	_, err = r.UpdateReactionHandler.Handle(command)
	if err != nil {
		log.Println("Failed to handle update command:", err)
		r.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	log.Println("Reaction updated successfully")
	r.Respond(c, http.StatusNoContent, gin.H{})
}


func (r ReactionController) DeleteReaction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		r.Respond(c, http.StatusBadRequest, gin.H{"error": "Id is not valid"})
		return
	}

	_, err = r.DeleteReactionHandler.Handle(id)

	if err != nil {
		r.Respond(c, http.StatusNotFound, er.NewBadRequest(err.Error()))
		return
	}

	r.Respond(c, http.StatusNoContent, gin.H{})

}
