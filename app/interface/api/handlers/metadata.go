package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	"github.com/polymorph-metadata/app/config"
	"github.com/polymorph-metadata/app/contracts"
	"github.com/polymorph-metadata/app/domain/metadata"
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	"github.com/polymorph-metadata/structs"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func HandleMetadataRequest(ethClient *ethereum.EthereumClient, polygonClient *ethereum.EthereumClient, address string, addressPolygon string, configService *config.ConfigService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		instanceRoot, err := contracts.NewPolymorphicFacesRoot(common.HexToAddress(address), ethClient.Client)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		instanceChild, err := contracts.NewPolymorphicFacesChild(common.HexToAddress(addressPolygon), polygonClient.Client)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		rootTunnelAddress := os.Getenv("ROOT_TUNNEL_ADDRESS")

		tokenId := r.URL.Query().Get("id")

		iTokenId, err := strconv.Atoi(tokenId)

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err)
			log.Errorln(err)
			return
		}

		ownerOf, err := instanceRoot.OwnerOf(nil, big.NewInt(int64(iTokenId)))

		var genomeInt *big.Int

		if ownerOf == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			msg := "Query for non-existing token"
			render.Status(r, 404)
			render.JSON(w, r, msg)
			log.Errorln(msg)
			return
		} else if ownerOf == common.HexToAddress(rootTunnelAddress) {
			genomeInt, err = instanceChild.GeneOf(nil, big.NewInt(int64(iTokenId)))
			if err != nil {
				render.Status(r, 500)
				render.JSON(w, r, err)
				log.Errorln(err)
				return
			}
		} else {
			genomeInt, err = instanceRoot.GeneOf(nil, big.NewInt(int64(iTokenId)))
			if err != nil {
				render.Status(r, 500)
				render.JSON(w, r, err)
				log.Errorln(err)
				return
			}
		}

		rarityResponse := structs.RarityServiceResponse{RarityScore: 0, Rank: 0}

		g := metadata.Genome(genomeInt.String())
		render.JSON(w, r, (&g).Metadata(tokenId, configService, rarityResponse))
	}
}
