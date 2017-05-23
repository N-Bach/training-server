package controller

type Controller struct {
	UserRepo IUserRepo
	LessonRepo ILessonRepo
	FeedbackRepo IFeedbackRepo
	ReviewRepo IReviewRepo
}

func NewController(userRepo IUserRepo, lessonRepo ILessonRepo, 
					feedbackRepo IFeedbackRepo, reviewRepo IReviewRepo) *Controller{
	return &Controller{
		UserRepo: userRepo,
		LessonRepo: lessonRepo,
		FeedbackRepo: feedbackRepo,
		ReviewRepo: reviewRepo,
	}
}

