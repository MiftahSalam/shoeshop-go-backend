package product

import "strings"

func searchFilter(keyword string) (*strings.Builder, []interface{}) {
	var (
		baseCond strings.Builder
		// query strings.Builder
		arg []interface{}
	)

	baseCond.WriteString("1=1")
	// if len(req.Filter) > 0 {
	// 	for i := range req.Filter {
	// 		split := strings.Split(req.Filter[i], "=")
	// 		if len(split) < 2 {
	// 			continue
	// 		}

	// 		switch split[0] {
	// 		case "start_date":
	// 			query.WriteString(" AND created_date >= ?")
	// 			arg = append(arg, split[1])
	// 		case "end_date":
	// 			query.WriteString(" AND created_date <= ?")
	// 			arg = append(arg, split[1])
	// 		case "partner_rrn":
	// 			query.WriteString(" AND partner_rrn = ?")
	// 			arg = append(arg, split[1])
	// 		case "channel_id":
	// 			query.WriteString(" AND channel_id = ?")
	// 			arg = append(arg, split[1])
	// 		case "trx_type":
	// 			query.WriteString(" AND trx_type = ?")
	// 			arg = append(arg, split[1])
	// 		case "trx_status":
	// 			query.WriteString(" AND trx_status = ?")
	// 			arg = append(arg, split[1])
	// 		}
	// 	}
	// 	baseCond.WriteString(query.String())
	// }

	if keyword != "" {
		baseCond.WriteString(" AND name LIKE ?")
		arg = append(arg, "%"+keyword+"%")
	}

	return &baseCond, arg

}
