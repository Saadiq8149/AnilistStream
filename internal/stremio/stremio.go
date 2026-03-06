package stremio

import (
	"anilist-stream/internal/anilist"
	"anilist-stream/internal/idmap"
	"anilist-stream/internal/metadata"
	"anilist-stream/internal/streams"
)

type StremioHandler struct {
	MetadataService *metadata.MetadataService
	SourceService   *streams.SourceService
	AnilistService  *anilist.AnilistService
	IDMapService    *idmap.IDMapService
}

func NewStremioHandler(metadataService *metadata.MetadataService, sourceService *streams.SourceService, anilistService *anilist.AnilistService, idMapService *idmap.IDMapService) *StremioHandler {
	return &StremioHandler{
		MetadataService: metadataService,
		SourceService:   sourceService,
		AnilistService:  anilistService,
		IDMapService:    idMapService,
	}
}
