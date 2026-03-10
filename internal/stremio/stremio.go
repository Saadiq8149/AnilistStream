package stremio

import (
	"anilist-stream/internal/anilist"
	"anilist-stream/internal/idmap"
	"anilist-stream/internal/metadata"
	"anilist-stream/internal/redis"
	"anilist-stream/internal/streams"
)

type StremioHandler struct {
	MetadataService *metadata.MetadataService
	SourceService   *streams.SourceService
	AnilistService  *anilist.AnilistService
	IDMapService    *idmap.IDMapService
	RedisService    *redis.RedisService
}

func NewStremioHandler(
	metadataService *metadata.MetadataService,
	sourceService *streams.SourceService,
	anilistService *anilist.AnilistService,
	idMapService *idmap.IDMapService,
	redisService *redis.RedisService) *StremioHandler {
	return &StremioHandler{
		MetadataService: metadataService,
		SourceService:   sourceService,
		AnilistService:  anilistService,
		IDMapService:    idMapService,
		RedisService:    redisService,
	}
}
