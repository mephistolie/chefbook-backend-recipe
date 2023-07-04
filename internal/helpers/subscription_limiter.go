package helpers

import (
	"github.com/mephistolie/chefbook-backend-common/subscription"
	"github.com/mephistolie/chefbook-backend-recipe/internal/config"
)

type SubscriptionLimiter struct {
	checkSubscription bool

	maxPicturesFree    int
	maxPicturesPremium int

	pictureMaxSizeFree    int64
	pictureMaxSizePremium int64
}

func NewSubscriptionLimiter(cfg config.Subscription) SubscriptionLimiter {
	return SubscriptionLimiter{
		checkSubscription:     *cfg.CheckSubscription,
		maxPicturesFree:       *cfg.MaxPicturesFree,
		maxPicturesPremium:    *cfg.MaxPicturesPremium,
		pictureMaxSizeFree:    *cfg.PictureMaxSizeFree,
		pictureMaxSizePremium: *cfg.PictureMaxSizePremium,
	}
}

func (l *SubscriptionLimiter) IsEncryptionAllowed(subscriptionPlan string) bool {
	return !l.checkSubscription || subscription.IsPremium(subscriptionPlan)
}

func (l *SubscriptionLimiter) GetMaxPicturesCount(subscriptionPlan string) int {
	if !l.checkSubscription || subscription.IsPremium(subscriptionPlan) {
		return l.maxPicturesPremium
	}
	return l.maxPicturesFree
}

func (l *SubscriptionLimiter) GetPictureMaxSize(subscriptionPlan string) int64 {
	if !l.checkSubscription || subscription.IsPremium(subscriptionPlan) {
		return l.pictureMaxSizePremium
	}
	return l.pictureMaxSizeFree
}
