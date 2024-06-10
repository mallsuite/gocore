package ml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type CloudService struct{}

func (s *CloudService) GetModuleTpl(serviceUserId int, serviceAppKey string) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	appRow := map[string]interface{}{}
	{
		goto LOOP_INIT_etzets
	LOOP_INIT_etzets:
		;

		d := 0
		goto LOOP_COND_uobqjq
	LOOP_COND_uobqjq:
		if d < 1 {
			goto LOOP_BODY_xkrbwe
		} else {
			goto LOOP_END_phzxut
		}
	LOOP_BODY_xkrbwe:
		{

			params := map[string]interface{}{
				"module_type":     2,
				"rtime":           time.Now().Unix(),
				"app_id_from":     100,
				"user_id_from":    serviceUserId,
				"service_app_key": serviceAppKey,
			}
			{
				goto LOOP_INIT_zyfbiy
			LOOP_INIT_zyfbiy:
				;

				d := 0
				goto LOOP_COND_vbnxhe
			LOOP_COND_vbnxhe:
				if d < 1 {
					goto LOOP_BODY_wkkqpd
				} else {
					goto LOOP_END_vxrcht
				}
			LOOP_BODY_wkkqpd:
				{
					d++
					goto LOOP_COND_vbnxhe
				}
			LOOP_END_vxrcht:
				{
				}
			}

			resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94103bc83302b940ed81586f4e2575bd5e42e54d9edc94019116bc024c36af54e5d026d1670020671f9078116b072"), params)
			if err != nil || resultStr == "" {
				resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978966211aab922099f1fc11cc5bfea461284e13453deb89b4d105c3cd938c24e6a96edad86c5020b42166254810fb6"), params)
				if err != nil {
					return nil, err
				}
			}

			var response map[string]interface{}
			{
				goto LOOP_INIT_chomqp
			LOOP_INIT_chomqp:
				;
				d := 0
				goto LOOP_COND_ktnmrw
			LOOP_COND_ktnmrw:
				if d < 1 {
					goto LOOP_BODY_rtbmjp
				} else {
					goto LOOP_END_pwqwmg
				}
			LOOP_BODY_rtbmjp:
				{
					err = json.Unmarshal([]byte(resultStr), &response)
					if err != nil {
						return nil, err
					}

					status := response["status"].(float64)
					if status == 250 {
						return nil, fmt.Errorf(response["msg"].(string))
					}
					d++
					goto LOOP_COND_ktnmrw

				}
			LOOP_END_pwqwmg:
				{
				}
			}

			appRow = response["data"].(map[string]interface{})
			d++
			goto LOOP_COND_uobqjq

		}
	LOOP_END_phzxut:
		{
		}
	}

	return appRow, nil
}

func (s *CloudService) Send(smsDto *SmsDto) error {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"sms_mobile":         smsDto.Mobile,
		"sms_content":        smsDto.Content,
		"tpl_id":             smsDto.TplId,
		"message_tpl_sender": smsDto.MessageTplSender,
		"tpl_paras":          smsDto.TplParas,
		"rtime":              time.Now().Unix(),
		"app_id_from":        100,
		"user_id_from":       smsDto.ServiceUserId,
		"service_app_key":    smsDto.ServiceAppKey,
	}
	{
		goto LOOP_INIT_ncrypv
	LOOP_INIT_ncrypv:
		;

		d := 0
		goto LOOP_COND_icbbdm
	LOOP_COND_icbbdm:
		if d < 1 {
			goto LOOP_BODY_toqccm
		} else {
			goto LOOP_END_kuhwlk
		}
	LOOP_BODY_toqccm:
		{
			d++
			goto LOOP_COND_icbbdm
		}
	LOOP_END_kuhwlk:
		{
		}
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e9500fbcc002038f57de1c8db6a9465698b52d54c2f0d66b1dc601112f621a8f4af516866210"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978876e11e98b0a12c619c81787f4fb4b5fd5e23448c3141bad6e3763572b6a430ac9b5861857"), params)
		if err != nil {
			return err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_lnzuan
	LOOP_INIT_lnzuan:
		;

		d := 0
		goto LOOP_COND_mpguia
	LOOP_COND_mpguia:
		if d < 1 {
			goto LOOP_BODY_foztla
		} else {
			goto LOOP_END_uabpol
		}
	LOOP_BODY_foztla:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return err
			}

			status := response["status"].(float64)
			if status == 250 {
				return fmt.Errorf(response["msg"].(string))
			}
			d++
			goto LOOP_COND_mpguia

		}
	LOOP_END_uabpol:
		{
		}
	}

	return nil
}

func (s *CloudService) ListSmsRecords(smsDto *SmsDto, page, rows int) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"rtime":           time.Now().Unix(),
		"app_id_from":     100,
		"user_id_from":    smsDto.ServiceUserId,
		"service_app_key": smsDto.ServiceAppKey,
		"page":            page,
		"rows":            rows,
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e9500fbcc002038f57c11090a6dd574c87fa23548bea96445d0b25c539b911b678b129391371f1da5cdb399bd4"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978876e11e98b0a12c606c40a9780ea51409aec3401d9e79f090a1239c4dd86a994fbceb9e4f1dc30bb494b55ee"), params)
		if err != nil {
			return nil, err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_ihuweq
	LOOP_INIT_ihuweq:
		;
		d := 0
		goto LOOP_COND_loiypq
	LOOP_COND_loiypq:
		if d < 1 {
			goto LOOP_BODY_mlffcr
		} else {
			goto LOOP_END_lgbpki
		}
	LOOP_BODY_mlffcr:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return nil, err
			}

			status := response["status"].(float64)
			if status == 250 {
				return nil, fmt.Errorf(response["msg"].(string))
			}

			return response["data"].(map[string]interface{}), nil
			d++
			goto LOOP_COND_loiypq

		}
	LOOP_END_lgbpki:
		{
		}
	}

	return nil, nil
}

func (s *CloudService) ListImOnline(serviceUserId int, serviceAppKey, puid string) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"rtime":           time.Now().Unix(),
		"app_id_from":     100,
		"user_id_from":    serviceUserId,
		"service_app_key": serviceAppKey,
		"puid":            puid,
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94a0f81891b0f980f8b1486a6b2554a9cdd3442dfd18158090f338c23d574719d3aa00e27c881387fc38bdaf0145b0bac887a37"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e2739789d6e2ca09206059e4cc01c97efe8575bbdfb2255e2f0835d0e0470de2edc39268426a1ec54d4b670dd8aa8b72af587417cf543"), params)
		if err != nil {
			return nil, err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_xmpzit
	LOOP_INIT_xmpzit:
		;

		d := 0
		goto LOOP_COND_boqcfm
	LOOP_COND_boqcfm:
		if d < 1 {
			goto LOOP_BODY_wgvkdo
		} else {
			goto LOOP_END_avxffa
		}
	LOOP_BODY_wgvkdo:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return nil, err
			}

			status := response["status"].(float64)
			if status == 250 {
				return nil, fmt.Errorf(response["msg"].(string))
			}

			return response["data"].(map[string]interface{}), nil
			d++
			goto LOOP_COND_boqcfm

		}
	LOOP_END_avxffa:
		{
		}
	}

	return nil, nil
}

func rtampgilxrizempotitnbxaaodcefbni(url string, params map[string]interface{}) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	reqURL := url + wxkwubclqqarkbdcpphdlwrgjeedzrqs(params)

	resp, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}
	{
		goto LOOP_INIT_hgtzbv
	LOOP_INIT_hgtzbv:
		;
		d := 0
		goto LOOP_COND_ludblg
	LOOP_COND_ludblg:
		if d < 1 {
			goto LOOP_BODY_hvcdhj
		} else {
			goto LOOP_END_dozlsh
		}
	LOOP_BODY_hvcdhj:
		{
			defer resp.Body.Close()
			d++
			goto LOOP_COND_ludblg

		}
	LOOP_END_dozlsh:
		{
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

func wxkwubclqqarkbdcpphdlwrgjeedzrqs(params map[string]interface{}) string {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	var queryString string

	for key, value := range params {
		switch value.(type) {
		case string:
			queryString += fmt.Sprintf(AES_DECRYPT("945cc9a74c81e58fa42fd015144f70a9b4993d2a8739"), key, url.QueryEscape(value.(string)))
		default:
			queryString += fmt.Sprintf(AES_DECRYPT("945cc9a74c81e58fa42fd015144f70a9b4993d2a8739"), key, value)
		}

	}
	return queryString
}

func (s *CloudService) InitService(params map[string]interface{}) error {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	{
		goto LOOP_INIT_tqvcog
	LOOP_INIT_tqvcog:
		;

		d := 0
		goto LOOP_COND_anwvbk
	LOOP_COND_anwvbk:
		if d < 1 {
			goto LOOP_BODY_vbqyhr
		} else {
			goto LOOP_END_kkrsdi
		}
	LOOP_BODY_vbqyhr:
		{
			d++
			goto LOOP_COND_anwvbk
		}
	LOOP_END_kkrsdi:
		{
		}
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94212a6b92e168b35f80987b3fb575dcee5225390fd8751030a70de2edc39268426a1468070f95fa377e18efe5adcabbaf70d9014b93e2ce8a34d89cb86777c6b769d401c80fc56b6e5f88d9084f7"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e27397895730b90a71f16a43fdd1d82a6ea400985ed331acef68a570b4722d327916e3f9827e906896de653b63ebb86b71fd3aeb6fc17d94eb17766e5ac44d3918e2d5b290caa05cbea74eff164878ecd2c"), params)
		if err != nil {
			return err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_wxmbhy
	LOOP_INIT_wxmbhy:
		;
		d := 0
		goto LOOP_COND_ynshjk
	LOOP_COND_ynshjk:
		if d < 1 {
			goto LOOP_BODY_pqjgmz
		} else {
			goto LOOP_END_zlsmqn
		}
	LOOP_BODY_pqjgmz:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return err
			}

			status := response["status"].(float64)
			if status == 250 {
				return fmt.Errorf(response["msg"].(string))
			}
			d++
			goto LOOP_COND_ynshjk

		}
	LOOP_END_zlsmqn:
		{
		}
	}

	return nil
}

func (s *CloudService) GetWeb(data *ReleaseDto) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	return s.GetRelease(data, AES_DECRYPT("c51cd8a18d7d7f6ddc026fbacdb9c603e776ae"))
}

func (s *CloudService) GetMpWeixin(data *ReleaseDto) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	return s.GetRelease(data, AES_DECRYPT("df0997ed0c9e7d9a84c83e93b8846b03535623e0230f5b3261"))
}

func (s *CloudService) GetMpApp(data *ReleaseDto) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	return s.GetRelease(data, AES_DECRYPT("d309cad9b6337e6721388bd54ecea68e63e39d"))
}

func (s *CloudService) GetRelease(data *ReleaseDto, releaseType string) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"rtime":           time.Now().Unix(),
		"release_type":    releaseType,
		"app_id_from":     data.ServiceAppId,
		"user_id_from":    data.ServiceUserId,
		"service_app_key": data.ServiceAppKey,
		"version":         data.Version,
		"url":             data.Url,
		"mp_app_id":       data.AppId,
		"mp_app_name":     data.AppName,
		"primary_color":   data.PrimaryColor,
	}
	{
		goto LOOP_INIT_ybavft
	LOOP_INIT_ybavft:
		;

		d := 0
		goto LOOP_COND_bmdgnh
	LOOP_COND_bmdgnh:
		if d < 1 {
			goto LOOP_BODY_tdvrys
		} else {
			goto LOOP_END_zyejpm
		}
	LOOP_BODY_tdvrys:
		{
			d++
			goto LOOP_COND_bmdgnh
		}
	LOOP_END_zyejpm:
		{
		}
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e95303ac8d0e019e29c21d86f4e2575bd5fa224bc8ff9c5146152fda6ac6772399408dc8cd5376115fd8cbbfb269b246d3"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978846201a4870803b805c91cc5bfea46129aed2b42cced8a12141826973ddf6b22641b2105b23b1e19908615060c3afbed"), params)
		if err != nil {
			return "", err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_jyizxr
	LOOP_INIT_jyizxr:
		;
		d := 0
		goto LOOP_COND_tbwjwv
	LOOP_COND_tbwjwv:
		if d < 1 {
			goto LOOP_BODY_ngalcn
		} else {
			goto LOOP_END_dpnrha
		}
	LOOP_BODY_ngalcn:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return "", err
			}

			status := response["status"].(float64)
			if status == 250 {
				return "", fmt.Errorf(response["msg"].(string))
			}
			d++
			goto LOOP_COND_tbwjwv

		}
	LOOP_END_dpnrha:
		{
		}
	}

	return response["data"].(string), nil
}

func (s *CloudService) GetDistrict(serviceUserId int, serviceAppKey string) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"rtime":           time.Now().Unix(),
		"app_id_from":     100,
		"user_id_from":    serviceUserId,
		"service_app_key": serviceAppKey,
	}
	{
		goto LOOP_INIT_crmxgf
	LOOP_INIT_crmxgf:
		;

		d := 0
		goto LOOP_COND_dsmxdl
	LOOP_COND_dsmxdl:
		if d < 1 {
			goto LOOP_BODY_tpxcyy
		} else {
			goto LOOP_END_gcvpnc
		}
	LOOP_BODY_tpxcyy:
		{
			d++
			goto LOOP_COND_dsmxdl
		}
	LOOP_END_gcvpnc:
		{
		}
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e9470bbc921d0f981e8b1486a6b25e469bfc3401d9e79f090a1239c4815398dc9118ca6683a29a06b2990bed"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978906a11bb9406058f4cc01c97efe35b5c9cfb6153d4eed25e130e388a610569454832ad3542ad0fe57a14d5"), params)
		if err != nil {
			return nil, err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_ayfwsu
	LOOP_INIT_ayfwsu:
		;
		d := 0
		goto LOOP_COND_lskync
	LOOP_COND_lskync:
		if d < 1 {
			goto LOOP_BODY_lyccrj
		} else {
			goto LOOP_END_czrbmq
		}
	LOOP_BODY_lyccrj:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return nil, err
			}

			status := response["status"].(float64)
			if status == 250 {
				return nil, fmt.Errorf(response["msg"].(string))
			}
			d++
			goto LOOP_COND_lskync

		}
	LOOP_END_czrbmq:
		{
		}
	}

	return response["data"].(map[string]interface{}), nil
}

func (s *CloudService) GetLicence(serviceUserId int, serviceAppKey string) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(1)
		{
			goto LOOP_INIT_elvcdq
		LOOP_INIT_elvcdq:
			;

			iXXX := 8
			goto LOOP_COND_dpptky
		LOOP_COND_dpptky:
			if iXXX < 15 {
				goto LOOP_BODY_izrpbn
			} else {
				goto LOOP_END_ppckbn
			}
		LOOP_BODY_izrpbn:
			{
				{
					goto LOOP_INIT_iykmte
				LOOP_INIT_iykmte:
					;
					jXXX := iXXX
					goto LOOP_COND_onmhah
				LOOP_COND_onmhah:
					if jXXX < 15 {
						goto LOOP_BODY_anoesz
					} else {
						goto LOOP_END_rzdvgc
					}
				LOOP_BODY_anoesz:
					{
						{
							goto LOOP_INIT_pmraeo
						LOOP_INIT_pmraeo:
							;
							zXXX := jXXX
							goto LOOP_COND_xvzdtl
						LOOP_COND_xvzdtl:
							if zXXX <
								15 {
								goto LOOP_BODY_fhahal
							} else {
								goto LOOP_END_cczqkr
							}
						LOOP_BODY_fhahal:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_xvzdtl
							}
						LOOP_END_cczqkr:
							{
							}
						}
						jXXX++
						goto LOOP_COND_onmhah

					}
				LOOP_END_rzdvgc:
					{
					}
				}
				iXXX++
				goto LOOP_COND_dpptky

			}
		LOOP_END_ppckbn:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	params := map[string]interface{}{
		"rtime":           time.Now().Unix(),
		"app_id_from":     100,
		"user_id_from":    serviceUserId,
		"service_app_key": serviceAppKey,
	}
	{
		goto LOOP_INIT_jnkzkk
	LOOP_INIT_jnkzkk:
		;

		d := 0
		goto LOOP_COND_pxvbeo
	LOOP_COND_pxvbeo:
		if d < 1 {
			goto LOOP_BODY_ipvnim
		} else {
			goto LOOP_END_smcvgu
		}
	LOOP_BODY_ipvnim:
		{
			d++
			goto LOOP_COND_pxvbeo
		}
	LOOP_END_smcvgu:
		{
		}
	}

	resultStr, err := rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94f0bac8301059e4cc01c97efe8575bcefc3e5790f49c5b0e4cbc9978abc5fd23b24dae714f6ca1f3"), params)
	if err != nil || resultStr == "" {
		resultStr, err = rtampgilxrizempotitnbxaaodcefbni(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978986a01aa880c03dd07c80ddeb5ea46099cf1371ac7ed805aed3f287fb196fac9cb450f458d8f939b"), params)
		if err != nil {
			return "", err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_rqifnj
	LOOP_INIT_rqifnj:
		;

		d := 0
		goto LOOP_COND_geuxog
	LOOP_COND_geuxog:
		if d < 1 {
			goto LOOP_BODY_fyidks
		} else {
			goto LOOP_END_vuidhr
		}
	LOOP_BODY_fyidks:
		{
			err = json.Unmarshal([]byte(resultStr), &response)
			if err != nil {
				return "", err
			}
			status := response["status"].(float64)
			if status == 250 {
				return "", fmt.Errorf(response["msg"].(string))
			}
			d++
			goto LOOP_COND_geuxog

		}
	LOOP_END_vuidhr:
		{
		}
	}

	return response["data"].(string), nil
}

type SmsDto struct {
	Mobile           string
	Content          string
	TplId            string
	MessageTplSender string
	TplParas         string
	ServiceUserId    int
	ServiceAppKey    string
}

type ReleaseDto struct {
	ServiceAppId  int
	ServiceUserId int
	ServiceAppKey string
	Version       string
	Url           string
	AppId         string
	AppName       string
	PrimaryColor  string
}
