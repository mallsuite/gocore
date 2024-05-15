package ml

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

func GetVer(ctx context.Context) string {
	(func() {
		zXXX := int64(2)
		sXXX := float64(10)
		{
			goto LOOP_INIT_slwppr
		LOOP_INIT_slwppr:
			;

			iXXX := 3
			goto LOOP_COND_clnqag
		LOOP_COND_clnqag:
			if iXXX < 15 {
				goto LOOP_BODY_eogpjy
			} else {
				goto LOOP_END_qsahef
			}
		LOOP_BODY_eogpjy:
			{
				{
					goto LOOP_INIT_cbwwpv
				LOOP_INIT_cbwwpv:
					;
					jXXX := iXXX
					goto LOOP_COND_psqqzf
				LOOP_COND_psqqzf:
					if jXXX < 15 {
						goto LOOP_BODY_dicbdw
					} else {
						goto LOOP_END_fciwdq
					}
				LOOP_BODY_dicbdw:
					{
						{
							goto LOOP_INIT_ouaoyr
						LOOP_INIT_ouaoyr:
							;
							zXXX := jXXX
							goto LOOP_COND_syeguo
						LOOP_COND_syeguo:
							if zXXX <
								15 {
								goto LOOP_BODY_lnpxvl
							} else {
								goto LOOP_END_rswrro
							}
						LOOP_BODY_lnpxvl:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_syeguo
							}
						LOOP_END_rswrro:
							{
							}
						}
						jXXX++
						goto LOOP_COND_psqqzf

					}
				LOOP_END_fciwdq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_clnqag

			}
		LOOP_END_qsahef:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	version, _ := g.Cfg().Get(ctx, "shopSuite.version")
	if version.String() == "true" {
		return version.String()
	} else {
		return ""
	}
	QueryExist := "false"
	{
		goto LOOP_INIT_xhvotu
	LOOP_INIT_xhvotu:
		;

		b := 0
		goto LOOP_COND_lxcpjs
	LOOP_COND_lxcpjs:
		if b < 1 && QueryExist != "false" {
			goto LOOP_BODY_kgzeql
		} else {
			goto LOOP_END_cxznkw
		}
	LOOP_BODY_kgzeql:
		{
			b++
			goto LOOP_COND_lxcpjs

		}
	LOOP_END_cxznkw:
		{
		}
	}
	{
		goto LOOP_INIT_pdwfus
	LOOP_INIT_pdwfus:
		;

		b := 0
		goto LOOP_COND_ztufht
	LOOP_COND_ztufht:
		if b < 1 {
			goto LOOP_BODY_gfrzhg
		} else {
			goto LOOP_END_nwolty
		}
	LOOP_BODY_gfrzhg:
		{
			b++
			goto LOOP_COND_ztufht
		}
	LOOP_END_nwolty:
		{
		}
	}
	{
		goto LOOP_INIT_bklfrr
	LOOP_INIT_bklfrr:
		;

		shs := 0
		goto LOOP_COND_klrzrz
	LOOP_COND_klrzrz:
		if shs < 1 {
			goto LOOP_BODY_ibxozk
		} else {
			goto LOOP_END_icmqwe
		}
	LOOP_BODY_ibxozk:
		{
			shs++
			goto LOOP_COND_klrzrz
		}
	LOOP_END_icmqwe:
		{
		}
	}
	{
		goto LOOP_INIT_snheql
	LOOP_INIT_snheql:
		;

		fsa := 0
		goto LOOP_COND_rlrlpq
	LOOP_COND_rlrlpq:
		if fsa < 1 {
			goto LOOP_BODY_lasmxl
		} else {
			goto LOOP_END_ahughv
		}
	LOOP_BODY_lasmxl:
		{
			fsa++
			goto LOOP_COND_rlrlpq
		}
	LOOP_END_ahughv:
		{
		}
	}

	return ""
}

func ConvertReqToInput(source, target interface{}, whereExt *[]*WhereExt) error {
	(func() {
		zXXX := int64(2)
		sXXX := float64(10)
		{
			goto LOOP_INIT_slwppr
		LOOP_INIT_slwppr:
			;

			iXXX := 3
			goto LOOP_COND_clnqag
		LOOP_COND_clnqag:
			if iXXX < 15 {
				goto LOOP_BODY_eogpjy
			} else {
				goto LOOP_END_qsahef
			}
		LOOP_BODY_eogpjy:
			{
				{
					goto LOOP_INIT_cbwwpv
				LOOP_INIT_cbwwpv:
					;
					jXXX := iXXX
					goto LOOP_COND_psqqzf
				LOOP_COND_psqqzf:
					if jXXX < 15 {
						goto LOOP_BODY_dicbdw
					} else {
						goto LOOP_END_fciwdq
					}
				LOOP_BODY_dicbdw:
					{
						{
							goto LOOP_INIT_ouaoyr
						LOOP_INIT_ouaoyr:
							;
							zXXX := jXXX
							goto LOOP_COND_syeguo
						LOOP_COND_syeguo:
							if zXXX <
								15 {
								goto LOOP_BODY_lnpxvl
							} else {
								goto LOOP_END_rswrro
							}
						LOOP_BODY_lnpxvl:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_syeguo
							}
						LOOP_END_rswrro:
							{
							}
						}
						jXXX++
						goto LOOP_COND_psqqzf

					}
				LOOP_END_fciwdq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_clnqag

			}
		LOOP_END_qsahef:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	return ConvertReqToInputWhere(source, target, whereExt)
}

func ConvertReqToInputWhere(source, target interface{}, whereExt *[]*WhereExt) error {
	(func() {
		zXXX := int64(2)
		sXXX := float64(10)
		{
			goto LOOP_INIT_slwppr
		LOOP_INIT_slwppr:
			;

			iXXX := 3
			goto LOOP_COND_clnqag
		LOOP_COND_clnqag:
			if iXXX < 15 {
				goto LOOP_BODY_eogpjy
			} else {
				goto LOOP_END_qsahef
			}
		LOOP_BODY_eogpjy:
			{
				{
					goto LOOP_INIT_cbwwpv
				LOOP_INIT_cbwwpv:
					;
					jXXX := iXXX
					goto LOOP_COND_psqqzf
				LOOP_COND_psqqzf:
					if jXXX < 15 {
						goto LOOP_BODY_dicbdw
					} else {
						goto LOOP_END_fciwdq
					}
				LOOP_BODY_dicbdw:
					{
						{
							goto LOOP_INIT_ouaoyr
						LOOP_INIT_ouaoyr:
							;
							zXXX := jXXX
							goto LOOP_COND_syeguo
						LOOP_COND_syeguo:
							if zXXX <
								15 {
								goto LOOP_BODY_lnpxvl
							} else {
								goto LOOP_END_rswrro
							}
						LOOP_BODY_lnpxvl:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_syeguo
							}
						LOOP_END_rswrro:
							{
							}
						}
						jXXX++
						goto LOOP_COND_psqqzf

					}
				LOOP_END_fciwdq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_clnqag

			}
		LOOP_END_qsahef:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()
	{
		goto LOOP_INIT_tcsiqs
	LOOP_INIT_tcsiqs:
		;

		a := 0
		goto LOOP_COND_irijdy
	LOOP_COND_irijdy:
		if a < 1 {
			goto LOOP_BODY_lxqeya
		} else {
			goto LOOP_END_jultob
		}
	LOOP_BODY_lxqeya:
		{
			sourceValue := reflect.ValueOf(source)
			if sourceValue.Kind() != reflect.Ptr || sourceValue.IsNil() {
				return errors.New("source must be a non-nil pointer")
			}

			targetValue := reflect.ValueOf(target)
			if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
				return errors.New("target must be a non-nil pointer")
			}

			sourceElem := sourceValue.Elem()
			targetElem := targetValue.Elem()

			if sourceElem.Type().Kind() != reflect.Struct || targetElem.Type().Kind() != reflect.Struct {
				return errors.New("source and target must be structs")
			}

			sourceType := sourceElem.Type()
			targetType := targetElem.Type()

			fieldS := sourceType.NumField()
			fieldT := targetType.NumField()

			fmt.Println(fieldS)
			fmt.Println(fieldT)
			{
				goto LOOP_INIT_acqpyr
			LOOP_INIT_acqpyr:
				;

				i := 0
				goto LOOP_COND_tagzhg
			LOOP_COND_tagzhg:
				if i < sourceType.NumField() {
					goto LOOP_BODY_vyrbcm
				} else {
					goto LOOP_END_jwedrs
				}
			LOOP_BODY_vyrbcm:
				{
					sourceField := sourceType.Field(i)
					sourceFieldValue := sourceElem.Field(i)
					columnName := sourceField.Name

					QueryField := sourceField.Tag.Get("field")
					QueryType := sourceField.Tag.Get("type")
					QueryExist := sourceField.Tag.Get("exist")
					{
						goto LOOP_INIT_kpaqbp
					LOOP_INIT_kpaqbp:
						;

						b := 0
						goto LOOP_COND_whxkbo
					LOOP_COND_whxkbo:
						if b < 1 && QueryExist != "false" {
							goto LOOP_BODY_somwpv
						} else {
							goto LOOP_END_mmozej
						}
					LOOP_BODY_somwpv:
						{
							if QueryField == "" {
							} else {
								columnName = gstr.CaseCamel(QueryField)
							}
							{
								goto LOOP_INIT_wnyovr
							LOOP_INIT_wnyovr:
								;

								t := 0
								goto LOOP_COND_yplycp
							LOOP_COND_yplycp:
								if t < targetType.NumField() {
									goto LOOP_BODY_njfrqj
								} else {
									goto LOOP_END_bgeyai
								}
							LOOP_BODY_njfrqj:
								{
									targetField := targetType.Field(t)

									if columnName == targetField.Name && sourceField.Name != "Meta" {
										if sourceFieldValue.IsValid() && !sourceFieldValue.IsZero() {
											{
												goto LOOP_INIT_zuidbo
											LOOP_INIT_zuidbo:
												;
												c := 0
												goto LOOP_COND_mgonfg
											LOOP_COND_mgonfg:
												if c < 1 && QueryType == "" {
													goto LOOP_BODY_dsekvj
												} else {
													goto LOOP_END_byluhe
												}
											LOOP_BODY_dsekvj:
												{
													*whereExt = append(*whereExt, &WhereExt{
														Column: columnName,
														Val:    sourceFieldValue.Interface(),
														Symbol: EQ,
													})
													c++
													goto LOOP_COND_mgonfg

												}
											LOOP_END_byluhe:
												{
												}
											}
											{
												goto LOOP_INIT_uunxth
											LOOP_INIT_uunxth:
												;

												c := 0
												goto LOOP_COND_hzklpo
											LOOP_COND_hzklpo:
												if c < 1 && QueryType != "" {
													goto LOOP_BODY_tecbgl
												} else {
													goto LOOP_END_fjdqgu
												}
											LOOP_BODY_tecbgl:
												{
													switch QueryType {
													case "EQ":
														{
															goto LOOP_INIT_srpqqd
														LOOP_INIT_srpqqd:
															;
															d := 0
															goto LOOP_COND_hfwywm
														LOOP_COND_hfwywm:
															if d < 1 {
																goto LOOP_BODY_wniwmq
															} else {
																goto LOOP_END_fnkrsa
															}
														LOOP_BODY_wniwmq:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: EQ,
																})
																d++
																goto LOOP_COND_hfwywm

															}
														LOOP_END_fnkrsa:
															{
															}
														}
													case "NE":
														{
															goto LOOP_INIT_oagpbk
														LOOP_INIT_oagpbk:
															;
															d := 0
															goto LOOP_COND_xokwzl
														LOOP_COND_xokwzl:
															if d < 1 {
																goto LOOP_BODY_hwgyoc
															} else {
																goto LOOP_END_vpjsxj
															}
														LOOP_BODY_hwgyoc:
															{
																d++
																goto LOOP_COND_xokwzl
															}
														LOOP_END_vpjsxj:
															{
															}
														}
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    sourceFieldValue.Interface(),
															Symbol: NE,
														})
													case "GT":
														{
															goto LOOP_INIT_fjgcgg
														LOOP_INIT_fjgcgg:
															;
															d := 0
															goto LOOP_COND_voanzm
														LOOP_COND_voanzm:
															if d < 1 {
																goto LOOP_BODY_akmrlt
															} else {
																goto LOOP_END_nmcljx
															}
														LOOP_BODY_akmrlt:
															{
																d++
																goto LOOP_COND_voanzm
															}
														LOOP_END_nmcljx:
															{
															}
														}
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    sourceFieldValue.Interface(),
															Symbol: GT,
														})
													case "GE":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    sourceFieldValue.Interface(),
															Symbol: GE,
														})
													case "LT":
														{
															goto LOOP_INIT_ymlxut
														LOOP_INIT_ymlxut:
															;
															d := 0
															goto LOOP_COND_ckudzl
														LOOP_COND_ckudzl:
															if d < 1 {
																goto LOOP_BODY_xjrosl
															} else {
																goto LOOP_END_qeoxyv
															}
														LOOP_BODY_xjrosl:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: LT,
																})
																d++
																goto LOOP_COND_ckudzl

															}
														LOOP_END_qeoxyv:
															{
															}
														}
													case "LE":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    sourceFieldValue.Interface(),
															Symbol: LE,
														})
													case "LIKE":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    "%" + sourceFieldValue.String() + "%",
															Symbol: LIKE,
														})
													case "NOT_LIKE":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    "%" + sourceFieldValue.String() + "%",
															Symbol: NOT_LIKE,
														})
													case "LIKE_LEFT":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    sourceFieldValue.String() + "%",
															Symbol: LIKE_LEFT,
														})
													case "LIKE_RIGHT":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    "%" + sourceFieldValue.String(),
															Symbol: LIKE_RIGHT,
														})
													case "IS_NULL":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    nil,
															Symbol: IS_NULL,
														})
													case "IS_NOT_NULL":
														*whereExt = append(*whereExt, &WhereExt{
															Column: columnName,
															Val:    nil,
															Symbol: IS_NOT_NULL,
														})
													case "IN":
														{
															goto LOOP_INIT_bseeve
														LOOP_INIT_bseeve:
															;
															d := 0
															goto LOOP_COND_edjaia
														LOOP_COND_edjaia:
															if d < 1 {
																goto LOOP_BODY_fjbink
															} else {
																goto LOOP_END_qdahwx
															}
														LOOP_BODY_fjbink:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN,
																})
																d++
																goto LOOP_COND_edjaia

															}
														LOOP_END_qdahwx:
															{
															}
														}
													case "NOT_IN":
														{
															goto LOOP_INIT_vvsjdv
														LOOP_INIT_vvsjdv:
															;
															d := 0
															goto LOOP_COND_zxungb
														LOOP_COND_zxungb:
															if d < 1 {
																goto LOOP_BODY_vzjfmi
															} else {
																goto LOOP_END_yeoqey
															}
														LOOP_BODY_vzjfmi:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN,
																})
																d++
																goto LOOP_COND_zxungb

															}
														LOOP_END_yeoqey:
															{
															}
														}
													case "IN_STR":
														{
															goto LOOP_INIT_yxifbl
														LOOP_INIT_yxifbl:
															;
															d := 0
															goto LOOP_COND_eusqte
														LOOP_COND_eusqte:
															if d < 1 {
																goto LOOP_BODY_ywobiz
															} else {
																goto LOOP_END_pnhpxm
															}
														LOOP_BODY_ywobiz:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN_STR,
																})
																d++
																goto LOOP_COND_eusqte

															}
														LOOP_END_pnhpxm:
															{
															}
														}
													case "NOT_IN_STR":
														{
															goto LOOP_INIT_bgmfco
														LOOP_INIT_bgmfco:
															;
															d := 0
															goto LOOP_COND_vtsxut
														LOOP_COND_vtsxut:
															if d < 1 {
																goto LOOP_BODY_royhgi
															} else {
																goto LOOP_END_hfsfzy
															}
														LOOP_BODY_royhgi:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN_STR,
																})
																d++
																goto LOOP_COND_vtsxut

															}
														LOOP_END_hfsfzy:
															{
															}
														}
													case "FIND_IN_SET":
														{
															goto LOOP_INIT_welfxj
														LOOP_INIT_welfxj:
															;
															d := 0
															goto LOOP_COND_yyjlnu
														LOOP_COND_yyjlnu:
															if d < 1 {
																goto LOOP_BODY_cbfdzs
															} else {
																goto LOOP_END_yiesxa
															}
														LOOP_BODY_cbfdzs:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET,
																})
																d++
																goto LOOP_COND_yyjlnu

															}
														LOOP_END_yiesxa:
															{
															}
														}
													case "FIND_IN_SET_STR":
														{
															goto LOOP_INIT_gfzwjd
														LOOP_INIT_gfzwjd:
															;
															d := 0
															goto LOOP_COND_spqpjm
														LOOP_COND_spqpjm:
															if d < 1 {
																goto LOOP_BODY_vbbvnx
															} else {
																goto LOOP_END_wfmuby
															}
														LOOP_BODY_vbbvnx:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET_STR,
																})
																d++
																goto LOOP_COND_spqpjm

															}
														LOOP_END_wfmuby:
															{
															}
														}
													default:
														panic(fmt.Sprintf("QueryType: %s is wrong", QueryType))
													}
													c++
													goto LOOP_COND_hzklpo

												}
											LOOP_END_fjdqgu:
												{
												}
											}
										}

									}
									t++
									goto LOOP_COND_yplycp

								}
							LOOP_END_bgeyai:
								{
								}
							}
							b++
							goto LOOP_COND_whxkbo

						}
					LOOP_END_mmozej:
						{
						}
					}
					i++
					goto LOOP_COND_tagzhg

				}
			LOOP_END_jwedrs:
				{
				}
			}
			a++
			goto LOOP_COND_irijdy

		}
	LOOP_END_jultob:
		{
		}
	}

	return nil
}

func BuildWhere(query *gdb.Model, whereLike []*WhereExt, whereExt []*WhereExt) (out *gdb.Model) {
	(func() {
		zXXX := int64(2)
		sXXX := float64(10)
		{
			goto LOOP_INIT_slwppr
		LOOP_INIT_slwppr:
			;

			iXXX := 3
			goto LOOP_COND_clnqag
		LOOP_COND_clnqag:
			if iXXX < 15 {
				goto LOOP_BODY_eogpjy
			} else {
				goto LOOP_END_qsahef
			}
		LOOP_BODY_eogpjy:
			{
				{
					goto LOOP_INIT_cbwwpv
				LOOP_INIT_cbwwpv:
					;
					jXXX := iXXX
					goto LOOP_COND_psqqzf
				LOOP_COND_psqqzf:
					if jXXX < 15 {
						goto LOOP_BODY_dicbdw
					} else {
						goto LOOP_END_fciwdq
					}
				LOOP_BODY_dicbdw:
					{
						{
							goto LOOP_INIT_ouaoyr
						LOOP_INIT_ouaoyr:
							;
							zXXX := jXXX
							goto LOOP_COND_syeguo
						LOOP_COND_syeguo:
							if zXXX <
								15 {
								goto LOOP_BODY_lnpxvl
							} else {
								goto LOOP_END_rswrro
							}
						LOOP_BODY_lnpxvl:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_syeguo
							}
						LOOP_END_rswrro:
							{
							}
						}
						jXXX++
						goto LOOP_COND_psqqzf

					}
				LOOP_END_fciwdq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_clnqag

			}
		LOOP_END_qsahef:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	if len(whereLike) > 0 {
		for _, c := range whereLike {
			query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
		}
	}

	if len(whereExt) > 0 {
		for _, c := range whereExt {
			switch c.Symbol {
			case EQ:
				{
					goto LOOP_INIT_towrxv
				LOOP_INIT_towrxv:
					;
					d := 0
					goto LOOP_COND_rqcpwb
				LOOP_COND_rqcpwb:
					if d < 1 {
						goto LOOP_BODY_pvwtor
					} else {
						goto LOOP_END_ksvuif
					}
				LOOP_BODY_pvwtor:
					{
						d++
						goto LOOP_COND_rqcpwb
					}
				LOOP_END_ksvuif:
					{
					}
				}
				query = query.Where(gstr.CaseSnake(c.Column), c.Val)
			case NE:
				{
					goto LOOP_INIT_gccufz
				LOOP_INIT_gccufz:
					;
					d := 0
					goto LOOP_COND_tpfvbv
				LOOP_COND_tpfvbv:
					if d < 1 {
						goto LOOP_BODY_fogqjg
					} else {
						goto LOOP_END_ggmrdy
					}
				LOOP_BODY_fogqjg:
					{
						d++
						goto LOOP_COND_tpfvbv
					}
				LOOP_END_ggmrdy:
					{
					}
				}
				query = query.WhereNot(gstr.CaseSnake(c.Column), c.Val)
			case GT:
				{
					goto LOOP_INIT_znjmds
				LOOP_INIT_znjmds:
					;
					d := 0
					goto LOOP_COND_kyndzk
				LOOP_COND_kyndzk:
					if d < 1 {
						goto LOOP_BODY_bvpzwh
					} else {
						goto LOOP_END_yzdvkv
					}
				LOOP_BODY_bvpzwh:
					{
						d++
						goto LOOP_COND_kyndzk
					}
				LOOP_END_yzdvkv:
					{
					}
				}
				query = query.WhereGT(gstr.CaseSnake(c.Column), c.Val)
			case GE:
				query = query.WhereGTE(gstr.CaseSnake(c.Column), c.Val)
			case LT:
				{
					goto LOOP_INIT_ijmmim
				LOOP_INIT_ijmmim:
					;
					d := 0
					goto LOOP_COND_wvblgo
				LOOP_COND_wvblgo:
					if d < 1 {
						goto LOOP_BODY_vvfoen
					} else {
						goto LOOP_END_ffyerq
					}
				LOOP_BODY_vvfoen:
					{
						d++
						goto LOOP_COND_wvblgo
					}
				LOOP_END_ffyerq:
					{
					}
				}
				query = query.WhereLT(gstr.CaseSnake(c.Column), c.Val)
			case LE:
				query = query.WhereLTE(gstr.CaseSnake(c.Column), c.Val)
			case LIKE:
				{
					goto LOOP_INIT_omgzjc
				LOOP_INIT_omgzjc:
					;
					d := 0
					goto LOOP_COND_johnpp
				LOOP_COND_johnpp:
					if d < 1 {
						goto LOOP_BODY_yjttmp
					} else {
						goto LOOP_END_usdedt
					}
				LOOP_BODY_yjttmp:
					{
						query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
						d++
						goto LOOP_COND_johnpp

					}
				LOOP_END_usdedt:
					{
					}
				}
			case NOT_LIKE:
				{
					goto LOOP_INIT_xjtpdk
				LOOP_INIT_xjtpdk:
					;
					d := 0
					goto LOOP_COND_vyexhy
				LOOP_COND_vyexhy:
					if d < 1 {
						goto LOOP_BODY_bgbhss
					} else {
						goto LOOP_END_uvifnu
					}
				LOOP_BODY_bgbhss:
					{
						d++
						goto LOOP_COND_vyexhy
					}
				LOOP_END_uvifnu:
					{
					}
				}
				query = query.WhereNotLike(gstr.CaseSnake(c.Column), c.Val.(string))
			case LIKE_LEFT:
				{
					goto LOOP_INIT_wrjqtm
				LOOP_INIT_wrjqtm:
					;
					d := 0
					goto LOOP_COND_dncsdn
				LOOP_COND_dncsdn:
					if d < 1 {
						goto LOOP_BODY_dfrwcd
					} else {
						goto LOOP_END_hpgnbq
					}
				LOOP_BODY_dfrwcd:
					{
						d++
						goto LOOP_COND_dncsdn
					}
				LOOP_END_hpgnbq:
					{
					}
				}
				query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
			case LIKE_RIGHT:
				query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
			case IS_NULL:
				query = query.WhereNull(gstr.CaseSnake(c.Column))
			case IS_NOT_NULL:
				query = query.WhereNotNull(gstr.CaseSnake(c.Column))
			case IN:
				{
					goto LOOP_INIT_mzfppj
				LOOP_INIT_mzfppj:
					;
					d := 0
					goto LOOP_COND_qfjlcb
				LOOP_COND_qfjlcb:
					if d < 1 {
						goto LOOP_BODY_ylzytm
					} else {
						goto LOOP_END_sdivmp
					}
				LOOP_BODY_ylzytm:
					{
						d++
						goto LOOP_COND_qfjlcb
					}
				LOOP_END_sdivmp:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereIn(gstr.CaseSnake(c.Column), c.Val)
				}
			case NOT_IN:
				{
					goto LOOP_INIT_osedde
				LOOP_INIT_osedde:
					;
					d := 0
					goto LOOP_COND_vszile
				LOOP_COND_vszile:
					if d < 1 {
						goto LOOP_BODY_svzmfi
					} else {
						goto LOOP_END_ogwmav
					}
				LOOP_BODY_svzmfi:
					{
						d++
						goto LOOP_COND_vszile
					}
				LOOP_END_ogwmav:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereNotIn(gstr.CaseSnake(c.Column), c.Val)
				}
			case IN_STR:
				if !g.IsEmpty(c.Val) {
					query = query.WhereIn(gstr.CaseSnake(c.Column), gstr.Split(c.Val.(string), ","))
				}
			case NOT_IN_STR:
				{
					goto LOOP_INIT_thscmb
				LOOP_INIT_thscmb:
					;
					d := 0
					goto LOOP_COND_avcqme
				LOOP_COND_avcqme:
					if d < 1 {
						goto LOOP_BODY_yibjkp
					} else {
						goto LOOP_END_lzuuhg
					}
				LOOP_BODY_yibjkp:
					{
						d++
						goto LOOP_COND_avcqme
					}
				LOOP_END_lzuuhg:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereNotIn(gstr.CaseSnake(c.Column), gstr.Split(c.Val.(string), ","))
				}
			case FIND_IN_SET:
				{
					goto LOOP_INIT_flmmeu
				LOOP_INIT_flmmeu:
					;
					d := 0
					goto LOOP_COND_uzezwx
				LOOP_COND_uzezwx:
					if d < 1 {
						goto LOOP_BODY_qcezoz
					} else {
						goto LOOP_END_xhnocf
					}
				LOOP_BODY_qcezoz:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range c.Val.([]uint64) {
								findInSetList = append(findInSetList, fmt.Sprintf("FIND_IN_SET(%d, %s)", u, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, " OR ") + " )")
						}
						d++
						goto LOOP_COND_uzezwx

					}
				LOOP_END_xhnocf:
					{
					}
				}
			case FIND_IN_SET_STR:
				{
					goto LOOP_INIT_hzwbld
				LOOP_INIT_hzwbld:
					;
					d := 0
					goto LOOP_COND_nloufa
				LOOP_COND_nloufa:
					if d < 1 {
						goto LOOP_BODY_gueqaf
					} else {
						goto LOOP_END_gqqcso
					}
				LOOP_BODY_gueqaf:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range gconv.SliceUint64(gstr.Split(c.Val.(string), ",")) {
								findInSetList = append(findInSetList, fmt.Sprintf("FIND_IN_SET(\"%d\", %s)", u, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, " OR ") + " )")
						}
						d++
						goto LOOP_COND_nloufa

					}
				LOOP_END_gqqcso:
					{
					}
				}
			}
		}
	}

	return query
}

func BuildOrder(query *gdb.Model, sidx string, sort string) (out *gdb.Model) {
	(func() {
		zXXX := int64(2)
		sXXX := float64(10)
		{
			goto LOOP_INIT_slwppr
		LOOP_INIT_slwppr:
			;

			iXXX := 3
			goto LOOP_COND_clnqag
		LOOP_COND_clnqag:
			if iXXX < 15 {
				goto LOOP_BODY_eogpjy
			} else {
				goto LOOP_END_qsahef
			}
		LOOP_BODY_eogpjy:
			{
				{
					goto LOOP_INIT_cbwwpv
				LOOP_INIT_cbwwpv:
					;
					jXXX := iXXX
					goto LOOP_COND_psqqzf
				LOOP_COND_psqqzf:
					if jXXX < 15 {
						goto LOOP_BODY_dicbdw
					} else {
						goto LOOP_END_fciwdq
					}
				LOOP_BODY_dicbdw:
					{
						{
							goto LOOP_INIT_ouaoyr
						LOOP_INIT_ouaoyr:
							;
							zXXX := jXXX
							goto LOOP_COND_syeguo
						LOOP_COND_syeguo:
							if zXXX <
								15 {
								goto LOOP_BODY_lnpxvl
							} else {
								goto LOOP_END_rswrro
							}
						LOOP_BODY_lnpxvl:
							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_syeguo
							}
						LOOP_END_rswrro:
							{
							}
						}
						jXXX++
						goto LOOP_COND_psqqzf

					}
				LOOP_END_fciwdq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_clnqag

			}
		LOOP_END_qsahef:
			{
			}
		}
		if sXXX == float64(zXXX) {
		}
	})()

	if !g.IsEmpty(sidx) {
		{
			goto LOOP_INIT_rbdysa
		LOOP_INIT_rbdysa:
			;
			d := 0
			goto LOOP_COND_egvgqy
		LOOP_COND_egvgqy:
			if d < 1 {
				goto LOOP_BODY_mgohnh
			} else {
				goto LOOP_END_mqemfx
			}
		LOOP_BODY_mgohnh:
			{
				if gstr.Equal(sort, ORDER_BY_DESC) {
					{
						goto LOOP_INIT_cmmlgo
					LOOP_INIT_cmmlgo:
						;
						d := 0
						goto LOOP_COND_cftjql
					LOOP_COND_cftjql:
						if d < 1 {
							goto LOOP_BODY_asazov
						} else {
							goto LOOP_END_jcdghm
						}
					LOOP_BODY_asazov:
						{
							query = query.OrderDesc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_cftjql

						}
					LOOP_END_jcdghm:
						{
						}
					}
				} else {
					{
						goto LOOP_INIT_lximwm
					LOOP_INIT_lximwm:
						;
						d := 0
						goto LOOP_COND_ahrxjg
					LOOP_COND_ahrxjg:
						if d < 1 {
							goto LOOP_BODY_wpoghq
						} else {
							goto LOOP_END_zctxwl
						}
					LOOP_BODY_wpoghq:
						{
							query = query.OrderAsc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_ahrxjg

						}
					LOOP_END_zctxwl:
						{
						}
					}
				}
				d++
				goto LOOP_COND_egvgqy

			}
		LOOP_END_mqemfx:
			{
			}
		}
	}

	return query
}

func AES_DECRYPT(s string) string {
	key, _ := hex.DecodeString("0101010101010101010101010101010101010101010101010101010101010101")
	ciphertext, _ := hex.DecodeString(s)
	nonce, _ := hex.DecodeString("010101010101010101010101")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
