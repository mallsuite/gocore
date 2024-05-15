package ml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CloudService struct{}

func (s *CloudService) GetModuleTpl(serviceUserId int, serviceAppKey string) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	appRow := map[string]interface{}{}
	{
		goto LOOP_INIT_qknlqf
	LOOP_INIT_qknlqf:
		;

		d := 0
		goto LOOP_COND_ydiojt
	LOOP_COND_ydiojt:
		if d < 1 {
			goto LOOP_BODY_nawemj
		} else {
			goto LOOP_END_nmjplp
		}
	LOOP_BODY_nawemj:
		{

			params := map[string]interface{}{
				"module_type":     2,
				"rtime":           time.Now().Unix(),
				"app_id_from":     100,
				"user_id_from":    serviceUserId,
				"service_app_key": serviceAppKey,
			}
			{
				goto LOOP_INIT_eshgql
			LOOP_INIT_eshgql:
				;

				d := 0
				goto LOOP_COND_egmmup
			LOOP_COND_egmmup:
				if d < 1 {
					goto LOOP_BODY_xulljc
				} else {
					goto LOOP_END_vxxxrq
				}
			LOOP_BODY_xulljc:
				{
					d++
					goto LOOP_COND_egmmup
				}
			LOOP_END_vxxxrq:
				{
				}
			}

			resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94103bc83302b940ed81586f4e2575bd5e42e54d9edc94019116bc024c36af54e5d026d1670020671f9078116b072"), params)
			if err != nil || resultStr == "" {
				resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978966211aab922099f1fc11cc5bfea461284e13453deb89b4d105c3cd938c24e6a96edad86c5020b42166254810fb6"), params)
				if err != nil {
					return nil, err
				}
			}

			var response map[string]interface{}
			{
				goto LOOP_INIT_gxxopo
			LOOP_INIT_gxxopo:
				;
				d := 0
				goto LOOP_COND_oenzyw
			LOOP_COND_oenzyw:
				if d < 1 {
					goto LOOP_BODY_gjylmp
				} else {
					goto LOOP_END_gkztuo
				}
			LOOP_BODY_gjylmp:
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
					goto LOOP_COND_oenzyw

				}
			LOOP_END_gkztuo:
				{
				}
			}

			appRow = response["data"].(map[string]interface{})
			d++
			goto LOOP_COND_ydiojt

		}
	LOOP_END_nmjplp:
		{
		}
	}

	return appRow, nil
}

func (s *CloudService) Send(smsDto *SmsDto) error {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		goto LOOP_INIT_hhqwya
	LOOP_INIT_hhqwya:
		;

		d := 0
		goto LOOP_COND_krtrup
	LOOP_COND_krtrup:
		if d < 1 {
			goto LOOP_BODY_dfyydv
		} else {
			goto LOOP_END_xmuknu
		}
	LOOP_BODY_dfyydv:
		{
			d++
			goto LOOP_COND_krtrup
		}
	LOOP_END_xmuknu:
		{
		}
	}

	resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e9500fbcc002038f57de1c8db6a9465698b52d54c2f0d66b1dc601112f621a8f4af516866210"), params)
	if err != nil || resultStr == "" {
		resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978876e11e98b0a12c619c81787f4fb4b5fd5e23448c3141bad6e3763572b6a430ac9b5861857"), params)
		if err != nil {
			return err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_iyxldm
	LOOP_INIT_iyxldm:
		;

		d := 0
		goto LOOP_COND_easfkg
	LOOP_COND_easfkg:
		if d < 1 {
			goto LOOP_BODY_gmells
		} else {
			goto LOOP_END_wybdie
		}
	LOOP_BODY_gmells:
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
			goto LOOP_COND_easfkg

		}
	LOOP_END_wybdie:
		{
		}
	}

	return nil
}

func mpxvslaohfmjfmkroachymdxbztvffxv(url string, params map[string]interface{}) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	reqURL := url + mhgkzajdxetglumpqktyuwpfuwxdnpet(params)

	resp, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}
	{
		goto LOOP_INIT_sqltvu
	LOOP_INIT_sqltvu:
		;
		d := 0
		goto LOOP_COND_ezjmgg
	LOOP_COND_ezjmgg:
		if d < 1 {
			goto LOOP_BODY_kqbpxo
		} else {
			goto LOOP_END_ahkhjq
		}
	LOOP_BODY_kqbpxo:
		{
			defer resp.Body.Close()
			d++
			goto LOOP_COND_ezjmgg

		}
	LOOP_END_ahkhjq:
		{
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

func mhgkzajdxetglumpqktyuwpfuwxdnpet(params map[string]interface{}) string {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	var queryString string
	for key, value := range params {
		queryString += fmt.Sprintf(AES_DECRYPT("945cc9a74c81e58fa42fd015144f70a9b4993d2a8739"), key, value)

	}
	return queryString
}

func (s *CloudService) InitService(params map[string]interface{}) error {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	{
		goto LOOP_INIT_rhnpyn
	LOOP_INIT_rhnpyn:
		;

		d := 0
		goto LOOP_COND_qfxbew
	LOOP_COND_qfxbew:
		if d < 1 {
			goto LOOP_BODY_nzvcqm
		} else {
			goto LOOP_END_wwjvlw
		}
	LOOP_BODY_nzvcqm:
		{
			d++
			goto LOOP_COND_qfxbew
		}
	LOOP_END_wwjvlw:
		{
		}
	}

	resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94212a6b92e168b35f80987b3fb575dcee5225390fd8751030a70de2edc39268426a1468070f95fa377e18efe5adcabbaf70d9014b93e2ce8a34d89cb86777c6b769d401c80fc56b6e5f88d9084f7"), params)
	if err != nil || resultStr == "" {
		resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e27397895730b90a71f16a43fdd1d82a6ea400985ed331acef68a570b4722d327916e3f9827e906896de653b63ebb86b71fd3aeb6fc17d94eb17766e5ac44d3918e2d5b290caa05cbea74eff164878ecd2c"), params)
		if err != nil {
			return err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_cdjmcd
	LOOP_INIT_cdjmcd:
		;
		d := 0
		goto LOOP_COND_rrjpwj
	LOOP_COND_rrjpwj:
		if d < 1 {
			goto LOOP_BODY_llhskw
		} else {
			goto LOOP_END_tmuzdd
		}
	LOOP_BODY_llhskw:
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
			goto LOOP_COND_rrjpwj

		}
	LOOP_END_tmuzdd:
		{
		}
	}

	return nil
}

func (s *CloudService) GetWeb(data *ReleaseDto) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
	}
	{
		goto LOOP_INIT_hottoz
	LOOP_INIT_hottoz:
		;

		d := 0
		goto LOOP_COND_fxpdpv
	LOOP_COND_fxpdpv:
		if d < 1 {
			goto LOOP_BODY_xtdtel
		} else {
			goto LOOP_END_twoufx
		}
	LOOP_BODY_xtdtel:
		{
			d++
			goto LOOP_COND_fxpdpv
		}
	LOOP_END_twoufx:
		{
		}
	}

	resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e95303ac8d0e019e29c21d86f4e2575bd5fa224bc8ff9c5146152fda6ac6772399408dc8cd5376115fd8cbbfb269b246d3"), params)
	if err != nil || resultStr == "" {
		resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978846201a4870803b805c91cc5bfea46129aed2b42cced8a12141826973ddf6b22641b2105b23b1e19908615060c3afbed"), params)
		if err != nil {
			return "", err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_xkfpeo
	LOOP_INIT_xkfpeo:
		;
		d := 0
		goto LOOP_COND_cbanty
	LOOP_COND_cbanty:
		if d < 1 {
			goto LOOP_BODY_eicwng
		} else {
			goto LOOP_END_xeouay
		}
	LOOP_BODY_eicwng:
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
			goto LOOP_COND_cbanty

		}
	LOOP_END_xeouay:
		{
		}
	}

	return response["data"].(string), nil
}

func (s *CloudService) GetDistrict(serviceUserId int, serviceAppKey string) (map[string]interface{}, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		goto LOOP_INIT_avthki
	LOOP_INIT_avthki:
		;

		d := 0
		goto LOOP_COND_dkqxfa
	LOOP_COND_dkqxfa:
		if d < 1 {
			goto LOOP_BODY_yzxapt
		} else {
			goto LOOP_END_bhgafb
		}
	LOOP_BODY_yzxapt:
		{
			d++
			goto LOOP_COND_dkqxfa
		}
	LOOP_END_bhgafb:
		{
		}
	}

	resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e9470bbc921d0f981e8b1486a6b25e469bfc3401d9e79f090a1239c4815398dc9118ca6683a29a06b2990bed"), params)
	if err != nil || resultStr == "" {
		resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978906a11bb9406058f4cc01c97efe35b5c9cfb6153d4eed25e130e388a610569454832ad3542ad0fe57a14d5"), params)
		if err != nil {
			return nil, err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_eflgwx
	LOOP_INIT_eflgwx:
		;
		d := 0
		goto LOOP_COND_pjmeli
	LOOP_COND_pjmeli:
		if d < 1 {
			goto LOOP_BODY_nvfzqf
		} else {
			goto LOOP_END_cpujvo
		}
	LOOP_BODY_nvfzqf:
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
			goto LOOP_COND_pjmeli

		}
	LOOP_END_cpujvo:
		{
		}
	}

	return response["data"].(map[string]interface{}), nil
}

func (s *CloudService) GetLicence(serviceUserId int, serviceAppKey string) (string, error) {
	(func() {
		zXXX := int64(3)
		sXXX := float64(2)
		{
			goto LOOP_INIT_obhkzb
		LOOP_INIT_obhkzb:
			;

			iXXX := 3
			goto LOOP_COND_eegbct
		LOOP_COND_eegbct:
			if iXXX < 15 {
				goto LOOP_BODY_htwhxy
			} else {
				goto LOOP_END_jzoeng
			}
		LOOP_BODY_htwhxy:
			{
				{
					goto LOOP_INIT_tptnhf
				LOOP_INIT_tptnhf:
					;
					jXXX := iXXX
					goto LOOP_COND_jhcqzs
				LOOP_COND_jhcqzs:
					if jXXX < 15 {
						goto LOOP_BODY_iybleb
					} else {
						goto LOOP_END_yxnnsz
					}
				LOOP_BODY_iybleb:
					{
						{
							goto LOOP_INIT_xhcbdn
						LOOP_INIT_xhcbdn:
							;
							zXXX := jXXX
							goto LOOP_COND_mqzuet
						LOOP_COND_mqzuet:
							if zXXX <
								15 {
								goto LOOP_BODY_eduepm
							} else {
								goto LOOP_END_rafrmf
							}
						LOOP_BODY_eduepm:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_mqzuet
							}
						LOOP_END_rafrmf:
							{
							}
						}
						jXXX++
						goto LOOP_COND_jhcqzs

					}
				LOOP_END_yxnnsz:
					{
					}
				}
				iXXX++
				goto LOOP_COND_eegbct

			}
		LOOP_END_jzoeng:
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
		goto LOOP_INIT_yprtqx
	LOOP_INIT_yprtqx:
		;

		d := 0
		goto LOOP_COND_ovevaz
	LOOP_COND_ovevaz:
		if d < 1 {
			goto LOOP_BODY_rnreve
		} else {
			goto LOOP_END_zxpztq
		}
	LOOP_BODY_rnreve:
		{
			d++
			goto LOOP_COND_ovevaz
		}
	LOOP_END_zxpztq:
		{
		}
	}

	resultStr, err := mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea1acd2adc8bc66cf2db523a876e5c0174228a3216e823ef10879bcb48488cf8094ae8386100ead365ef3af8cfd74f1b302129e94f0bac8301059e4cc01c97efe8575bcefc3e5790f49c5b0e4cbc9978abc5fd23b24dae714f6ca1f3"), params)
	if err != nil || resultStr == "" {
		resultStr, err = mpxvslaohfmjfmkroachymdxbztvffxv(AES_DECRYPT("da0dceea53d82a9289c660e8c04860da755b1e7724962f07a36ee251c19cc14955daa61152a76a6811a29d73f83ee7c5d10c5e273978986a01aa880c03dd07c80ddeb5ea46099cf1371ac7ed805aed3f287fb196fac9cb450f458d8f939b"), params)
		if err != nil {
			return "", err
		}
	}

	var response map[string]interface{}
	{
		goto LOOP_INIT_nzpypq
	LOOP_INIT_nzpypq:
		;

		d := 0
		goto LOOP_COND_ledcfm
	LOOP_COND_ledcfm:
		if d < 1 {
			goto LOOP_BODY_jtsuwd
		} else {
			goto LOOP_END_yehftr
		}
	LOOP_BODY_jtsuwd:
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
			goto LOOP_COND_ledcfm

		}
	LOOP_END_yehftr:
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
}
