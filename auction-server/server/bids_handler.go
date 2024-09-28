package server

import (
	"net/http"

	"github.com/subham-singh46/auction/pkg/utils"
)

func (s *Server) AddNewBid(w http.ResponseWriter, r *http.Request) {
	d := &AddNewBidReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again later", http.StatusInternalServerError)
		return
	}

	userId := r.Context().Value("UserID").(int)
	bidId, err := s.store.AddNewBid(d.BidPrice, d.OwnerID, d.TicketID, userId)
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again later", http.StatusInternalServerError)
		return
	}

	res := &AddNewBidRes{BidID: bidId}
	utils.WriteResponse(w, nil, res, http.StatusOK)
	return
}

func (s *Server) GetUserBids(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("UserID").(int)
	bids, err := s.store.GetUserBids(userId)
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again later", http.StatusInternalServerError)
		return
	}

	res := &GetUserBidsRes{}
	resBids := make([]UserBid, 0)
	for _, b := range bids {
		bid := UserBid{
			BidId:         b.BidId,
			TicketId:      b.TicketId,
			Venue:         b.Venue,
			OriginalPrice: b.OriginalPrice,
			BidPrice:      b.BidPrice,
			OwnerId:       b.OwnerId,
			BidderId:      b.BidderId,
			CreatedAt:     b.CreatedAt.String(),
		}
		resBids = append(resBids, bid)
	}
	res.Bids = resBids
	utils.WriteResponse(w, nil, res, http.StatusOK)
	return
}
