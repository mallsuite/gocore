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
		zXXX := int64(1)
		sXXX := float64(6)
		{
			goto LOOP_INIT_dfzygz
		LOOP_INIT_dfzygz:
			;
			iXXX :=

				1
			goto LOOP_COND_exvnbb
		LOOP_COND_exvnbb:
			if iXXX < 15 {
				goto LOOP_BODY_visexe
			} else {
				goto LOOP_END_mcfgyg
			}
		LOOP_BODY_visexe:
			{
				{
					goto LOOP_INIT_stnpvn
				LOOP_INIT_stnpvn:
					;
					jXXX := iXXX
					goto LOOP_COND_ihefrv
				LOOP_COND_ihefrv:
					if jXXX < 15 {
						goto LOOP_BODY_uwxdmm
					} else {
						goto LOOP_END_akaffq
					}
				LOOP_BODY_uwxdmm:
					{
						{
							goto LOOP_INIT_wsjowg
						LOOP_INIT_wsjowg:
							;
							zXXX := jXXX
							goto LOOP_COND_yajcfy
						LOOP_COND_yajcfy:
							if zXXX < 15 {
								goto LOOP_BODY_zspagn
							} else {
								goto LOOP_END_usnuqw
							}
						LOOP_BODY_zspagn:

							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_yajcfy

							}
						LOOP_END_usnuqw:
							{
							}
						}
						jXXX++
						goto LOOP_COND_ihefrv

					}
				LOOP_END_akaffq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_exvnbb

			}
		LOOP_END_mcfgyg:
			{
			}
		}
		if sXXX ==
			float64(zXXX) {
		}
	})()
	version, _ := g.Cfg().Get(ctx, AES_DECRYPT("c111d5ea3a826c878f8b79f8dc4f27c673fc5ccad57aee767211263f912076e1ab"))
	if version.String() == "true" {
		return version.String()
	} else {
		return ""
	}
	QueryExist := AES_DECRYPT("d418d6e90c20c5570e003b02d3471bc053055f8df7")
	{
		goto LOOP_INIT_wzniwk
	LOOP_INIT_wzniwk:
		;

		b := 0
		goto LOOP_COND_knyekq
	LOOP_COND_knyekq:
		if b < 1 && QueryExist != "false" {
			goto LOOP_BODY_ibshtk
		} else {
			goto LOOP_END_wsxfxv
		}
	LOOP_BODY_ibshtk:
		{
			b++
			goto LOOP_COND_knyekq

		}
	LOOP_END_wsxfxv:
		{
		}
	}
	{
		goto LOOP_INIT_wfvwtn
	LOOP_INIT_wfvwtn:
		;

		b := 0
		goto LOOP_COND_qncyja
	LOOP_COND_qncyja:
		if b < 1 {
			goto LOOP_BODY_guuqcp
		} else {
			goto LOOP_END_uaqmmf
		}
	LOOP_BODY_guuqcp:
		{
			b++
			goto LOOP_COND_qncyja
		}
	LOOP_END_uaqmmf:
		{
		}
	}
	{
		goto LOOP_INIT_ybqfwt
	LOOP_INIT_ybqfwt:
		;

		shs := 0
		goto LOOP_COND_odwlod
	LOOP_COND_odwlod:
		if shs < 1 {
			goto LOOP_BODY_kjbgjf
		} else {
			goto LOOP_END_cawlat
		}
	LOOP_BODY_kjbgjf:
		{
			shs++
			goto LOOP_COND_odwlod
		}
	LOOP_END_cawlat:
		{
		}
	}
	{
		goto LOOP_INIT_vvcjzi
	LOOP_INIT_vvcjzi:
		;

		fsa := 0
		goto LOOP_COND_xeojab
	LOOP_COND_xeojab:
		if fsa < 1 {
			goto LOOP_BODY_esgeqa
		} else {
			goto LOOP_END_esuxnp
		}
	LOOP_BODY_esgeqa:
		{
			fsa++
			goto LOOP_COND_xeojab
		}
	LOOP_END_esuxnp:
		{
		}
	}

	return ""
}

func ConvertReqToInput(source, target interface{}, whereExt *[]*WhereExt) error {
	(func() {
		zXXX := int64(1)
		sXXX := float64(6)
		{
			goto LOOP_INIT_dfzygz
		LOOP_INIT_dfzygz:
			;
			iXXX :=

				1
			goto LOOP_COND_exvnbb
		LOOP_COND_exvnbb:
			if iXXX < 15 {
				goto LOOP_BODY_visexe
			} else {
				goto LOOP_END_mcfgyg
			}
		LOOP_BODY_visexe:
			{
				{
					goto LOOP_INIT_stnpvn
				LOOP_INIT_stnpvn:
					;
					jXXX := iXXX
					goto LOOP_COND_ihefrv
				LOOP_COND_ihefrv:
					if jXXX < 15 {
						goto LOOP_BODY_uwxdmm
					} else {
						goto LOOP_END_akaffq
					}
				LOOP_BODY_uwxdmm:
					{
						{
							goto LOOP_INIT_wsjowg
						LOOP_INIT_wsjowg:
							;
							zXXX := jXXX
							goto LOOP_COND_yajcfy
						LOOP_COND_yajcfy:
							if zXXX < 15 {
								goto LOOP_BODY_zspagn
							} else {
								goto LOOP_END_usnuqw
							}
						LOOP_BODY_zspagn:

							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_yajcfy

							}
						LOOP_END_usnuqw:
							{
							}
						}
						jXXX++
						goto LOOP_COND_ihefrv

					}
				LOOP_END_akaffq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_exvnbb

			}
		LOOP_END_mcfgyg:
			{
			}
		}
		if sXXX ==
			float64(zXXX) {
		}
	})()

	return ConvertReqToInputWhere(source, target, whereExt)
}

func ConvertReqToInputWhere(source, target interface{}, whereExt *[]*WhereExt) error {
	(func() {
		zXXX := int64(1)
		sXXX := float64(6)
		{
			goto LOOP_INIT_dfzygz
		LOOP_INIT_dfzygz:
			;
			iXXX :=

				1
			goto LOOP_COND_exvnbb
		LOOP_COND_exvnbb:
			if iXXX < 15 {
				goto LOOP_BODY_visexe
			} else {
				goto LOOP_END_mcfgyg
			}
		LOOP_BODY_visexe:
			{
				{
					goto LOOP_INIT_stnpvn
				LOOP_INIT_stnpvn:
					;
					jXXX := iXXX
					goto LOOP_COND_ihefrv
				LOOP_COND_ihefrv:
					if jXXX < 15 {
						goto LOOP_BODY_uwxdmm
					} else {
						goto LOOP_END_akaffq
					}
				LOOP_BODY_uwxdmm:
					{
						{
							goto LOOP_INIT_wsjowg
						LOOP_INIT_wsjowg:
							;
							zXXX := jXXX
							goto LOOP_COND_yajcfy
						LOOP_COND_yajcfy:
							if zXXX < 15 {
								goto LOOP_BODY_zspagn
							} else {
								goto LOOP_END_usnuqw
							}
						LOOP_BODY_zspagn:

							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_yajcfy

							}
						LOOP_END_usnuqw:
							{
							}
						}
						jXXX++
						goto LOOP_COND_ihefrv

					}
				LOOP_END_akaffq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_exvnbb

			}
		LOOP_END_mcfgyg:
			{
			}
		}
		if sXXX ==
			float64(zXXX) {
		}
	})()
	{
		goto LOOP_INIT_qphuln
	LOOP_INIT_qphuln:
		;

		a := 0
		goto LOOP_COND_yiygbg
	LOOP_COND_yiygbg:
		if a < 1 {
			goto LOOP_BODY_twlyal
		} else {
			goto LOOP_END_xwsetg
		}
	LOOP_BODY_twlyal:
		{
			sourceValue := reflect.ValueOf(source)
			if sourceValue.Kind() != reflect.Ptr || sourceValue.IsNil() {
				return errors.New(AES_DECRYPT("c116cfe80a92259e9fd67bbdcc596ec83d5a016a7c91320ead7de317c686c05ee3d4d58fc4bc21ccc4ffd5ff756616f2"))
			}

			targetValue := reflect.ValueOf(target)
			if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
				return errors.New(AES_DECRYPT("c618c8fd0c83259e9fd67bbdcc596ec83d5a016a7c91320ead7de317c686c05ebec56395ed40e13e53cd1b7969f4eb6a"))
			}

			sourceElem := sourceValue.Elem()
			targetElem := targetValue.Elem()

			if sourceElem.Type().Kind() != reflect.Struct || targetElem.Type().Kind() != reflect.Struct {
				return errors.New(AES_DECRYPT("c116cfe80a92259284c12fe9cf4e29cc69140371228b7b00e82dff0ada87c6585e96a2f147f30cfe034ef2d60e831a91bf"))
			}

			sourceType := sourceElem.Type()
			targetType := targetElem.Type()

			fieldS := sourceType.NumField()
			fieldT := targetType.NumField()

			fmt.Println(fieldS)
			fmt.Println(fieldT)
			{
				goto LOOP_INIT_dytcyx
			LOOP_INIT_dytcyx:
				;

				i := 0
				goto LOOP_COND_cqycms
			LOOP_COND_cqycms:
				if i < sourceType.NumField() {
					goto LOOP_BODY_vvuvyk
				} else {
					goto LOOP_END_sxzqjf
				}
			LOOP_BODY_vvuvyk:
				{
					sourceField := sourceType.Field(i)
					sourceFieldValue := sourceElem.Field(i)
					columnName := sourceField.Name

					QueryField := sourceField.Tag.Get(AES_DECRYPT("d410dff60dabbed5e8a40cf6c2b9b0fc4d111bf39d"))
					QueryType := sourceField.Tag.Get(AES_DECRYPT("c600caff9f942ede2e68e918fd4d2409bcaf9fff"))
					QueryExist := sourceField.Tag.Get(AES_DECRYPT("d701d3e91d36b505af3f4815f75c6417af9d56646a"))
					{
						goto LOOP_INIT_uqbewt
					LOOP_INIT_uqbewt:
						;

						b := 0
						goto LOOP_COND_cqrmle
					LOOP_COND_cqrmle:
						if b < 1 && QueryExist != "false" {
							goto LOOP_BODY_yfxsax
						} else {
							goto LOOP_END_mhhhko
						}
					LOOP_BODY_yfxsax:
						{
							if QueryField == "" {
							} else {
								columnName = gstr.CaseCamel(QueryField)
							}
							{
								goto LOOP_INIT_qdgsbx
							LOOP_INIT_qdgsbx:
								;

								t := 0
								goto LOOP_COND_gadjqw
							LOOP_COND_gadjqw:
								if t < targetType.NumField() {
									goto LOOP_BODY_wecsja
								} else {
									goto LOOP_END_wzfoan
								}
							LOOP_BODY_wecsja:
								{
									targetField := targetType.Field(t)

									if columnName == targetField.Name && sourceField.Name != "Meta" {
										if sourceFieldValue.IsValid() && !sourceFieldValue.IsZero() {
											{
												goto LOOP_INIT_dehqyd
											LOOP_INIT_dehqyd:
												;
												c := 0
												goto LOOP_COND_fkyoaf
											LOOP_COND_fkyoaf:
												if c < 1 && QueryType == "" {
													goto LOOP_BODY_hhtymg
												} else {
													goto LOOP_END_nplqoq
												}
											LOOP_BODY_hhtymg:
												{
													*whereExt = append(*whereExt, &WhereExt{
														Column: columnName,
														Val:    sourceFieldValue.Interface(),
														Symbol: EQ,
													})
													c++
													goto LOOP_COND_fkyoaf

												}
											LOOP_END_nplqoq:
												{
												}
											}
											{
												goto LOOP_INIT_ccxoxv
											LOOP_INIT_ccxoxv:
												;

												c := 0
												goto LOOP_COND_blzkrx
											LOOP_COND_blzkrx:
												if c < 1 && QueryType != "" {
													goto LOOP_BODY_vavzlv
												} else {
													goto LOOP_END_xwhpvg
												}
											LOOP_BODY_vavzlv:
												{
													switch QueryType {
													case "EQ":
														{
															goto LOOP_INIT_khbchl
														LOOP_INIT_khbchl:
															;
															d := 0
															goto LOOP_COND_srriyq
														LOOP_COND_srriyq:
															if d < 1 {
																goto LOOP_BODY_negevw
															} else {
																goto LOOP_END_znwdji
															}
														LOOP_BODY_negevw:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: EQ,
																})
																d++
																goto LOOP_COND_srriyq

															}
														LOOP_END_znwdji:
															{
															}
														}
													case "NE":
														{
															goto LOOP_INIT_mtliau
														LOOP_INIT_mtliau:
															;
															d := 0
															goto LOOP_COND_qqmhsy
														LOOP_COND_qqmhsy:
															if d < 1 {
																goto LOOP_BODY_xcwyzy
															} else {
																goto LOOP_END_tkpeag
															}
														LOOP_BODY_xcwyzy:
															{
																d++
																goto LOOP_COND_qqmhsy
															}
														LOOP_END_tkpeag:
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
															goto LOOP_INIT_lxkogh
														LOOP_INIT_lxkogh:
															;
															d := 0
															goto LOOP_COND_npkiwj
														LOOP_COND_npkiwj:
															if d < 1 {
																goto LOOP_BODY_pxbksj
															} else {
																goto LOOP_END_znmyom
															}
														LOOP_BODY_pxbksj:
															{
																d++
																goto LOOP_COND_npkiwj
															}
														LOOP_END_znmyom:
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
															goto LOOP_INIT_pcscmi
														LOOP_INIT_pcscmi:
															;
															d := 0
															goto LOOP_COND_pmsygi
														LOOP_COND_pmsygi:
															if d < 1 {
																goto LOOP_BODY_yciwrw
															} else {
																goto LOOP_END_rertwa
															}
														LOOP_BODY_yciwrw:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: LT,
																})
																d++
																goto LOOP_COND_pmsygi

															}
														LOOP_END_rertwa:
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
															goto LOOP_INIT_euyoks
														LOOP_INIT_euyoks:
															;
															d := 0
															goto LOOP_COND_bpkgdy
														LOOP_COND_bpkgdy:
															if d < 1 {
																goto LOOP_BODY_wgzdya
															} else {
																goto LOOP_END_adxpyx
															}
														LOOP_BODY_wgzdya:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN,
																})
																d++
																goto LOOP_COND_bpkgdy

															}
														LOOP_END_adxpyx:
															{
															}
														}
													case "NOT_IN":
														{
															goto LOOP_INIT_uuppxu
														LOOP_INIT_uuppxu:
															;
															d := 0
															goto LOOP_COND_igekvv
														LOOP_COND_igekvv:
															if d < 1 {
																goto LOOP_BODY_rvhmrg
															} else {
																goto LOOP_END_mkmkyd
															}
														LOOP_BODY_rvhmrg:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN,
																})
																d++
																goto LOOP_COND_igekvv

															}
														LOOP_END_mkmkyd:
															{
															}
														}
													case "IN_STR":
														{
															goto LOOP_INIT_hyzwnx
														LOOP_INIT_hyzwnx:
															;
															d := 0
															goto LOOP_COND_utkfjj
														LOOP_COND_utkfjj:
															if d < 1 {
																goto LOOP_BODY_wdfzjv
															} else {
																goto LOOP_END_srwpoo
															}
														LOOP_BODY_wdfzjv:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN_STR,
																})
																d++
																goto LOOP_COND_utkfjj

															}
														LOOP_END_srwpoo:
															{
															}
														}
													case "NOT_IN_STR":
														{
															goto LOOP_INIT_ozvrez
														LOOP_INIT_ozvrez:
															;
															d := 0
															goto LOOP_COND_vxvztr
														LOOP_COND_vxvztr:
															if d < 1 {
																goto LOOP_BODY_gftelm
															} else {
																goto LOOP_END_doayaa
															}
														LOOP_BODY_gftelm:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN_STR,
																})
																d++
																goto LOOP_COND_vxvztr

															}
														LOOP_END_doayaa:
															{
															}
														}
													case "FIND_IN_SET":
														{
															goto LOOP_INIT_umflpb
														LOOP_INIT_umflpb:
															;
															d := 0
															goto LOOP_COND_kyqbxt
														LOOP_COND_kyqbxt:
															if d < 1 {
																goto LOOP_BODY_ozqnrz
															} else {
																goto LOOP_END_pvuvvw
															}
														LOOP_BODY_ozqnrz:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET,
																})
																d++
																goto LOOP_COND_kyqbxt

															}
														LOOP_END_pvuvvw:
															{
															}
														}
													case "FIND_IN_SET_STR":
														{
															goto LOOP_INIT_nfccek
														LOOP_INIT_nfccek:
															;
															d := 0
															goto LOOP_COND_tyanzu
														LOOP_COND_tyanzu:
															if d < 1 {
																goto LOOP_BODY_ygqfse
															} else {
																goto LOOP_END_xduotx
															}
														LOOP_BODY_ygqfse:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET_STR,
																})
																d++
																goto LOOP_COND_tyanzu

															}
														LOOP_END_xduotx:
															{
															}
														}
													case "FIND_IN_SET_STR_STR":
														{
															goto LOOP_INIT_clojaz
														LOOP_INIT_clojaz:
															;
															d := 0
															goto LOOP_COND_fmtrre
														LOOP_COND_fmtrre:
															if d < 1 {
																goto LOOP_BODY_igvsrf
															} else {
																goto LOOP_END_gyyglm
															}
														LOOP_BODY_igvsrf:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET_STR_STR,
																})
																d++
																goto LOOP_COND_fmtrre

															}
														LOOP_END_gyyglm:
															{
															}
														}
													default:
														panic(fmt.Sprintf(AES_DECRYPT("e30cdfe810a37c838f9f2fb8dd1c27da3d431c6b3f98d6d8dabc18f301aa0e0b16cce4736c90"), QueryType))
													}
													c++
													goto LOOP_COND_blzkrx

												}
											LOOP_END_xwhpvg:
												{
												}
											}
										}

									}
									t++
									goto LOOP_COND_gadjqw

								}
							LOOP_END_wzfoan:
								{
								}
							}
							b++
							goto LOOP_COND_cqrmle

						}
					LOOP_END_mhhhko:
						{
						}
					}
					i++
					goto LOOP_COND_cqycms

				}
			LOOP_END_sxzqjf:
				{
				}
			}
			a++
			goto LOOP_COND_yiygbg

		}
	LOOP_END_xwsetg:
		{
		}
	}

	return nil
}

func BuildWhere(query *gdb.Model, whereLike []*WhereExt, whereExt []*WhereExt) (out *gdb.Model) {
	(func() {
		zXXX := int64(1)
		sXXX := float64(6)
		{
			goto LOOP_INIT_dfzygz
		LOOP_INIT_dfzygz:
			;
			iXXX :=

				1
			goto LOOP_COND_exvnbb
		LOOP_COND_exvnbb:
			if iXXX < 15 {
				goto LOOP_BODY_visexe
			} else {
				goto LOOP_END_mcfgyg
			}
		LOOP_BODY_visexe:
			{
				{
					goto LOOP_INIT_stnpvn
				LOOP_INIT_stnpvn:
					;
					jXXX := iXXX
					goto LOOP_COND_ihefrv
				LOOP_COND_ihefrv:
					if jXXX < 15 {
						goto LOOP_BODY_uwxdmm
					} else {
						goto LOOP_END_akaffq
					}
				LOOP_BODY_uwxdmm:
					{
						{
							goto LOOP_INIT_wsjowg
						LOOP_INIT_wsjowg:
							;
							zXXX := jXXX
							goto LOOP_COND_yajcfy
						LOOP_COND_yajcfy:
							if zXXX < 15 {
								goto LOOP_BODY_zspagn
							} else {
								goto LOOP_END_usnuqw
							}
						LOOP_BODY_zspagn:

							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_yajcfy

							}
						LOOP_END_usnuqw:
							{
							}
						}
						jXXX++
						goto LOOP_COND_ihefrv

					}
				LOOP_END_akaffq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_exvnbb

			}
		LOOP_END_mcfgyg:
			{
			}
		}
		if sXXX ==
			float64(zXXX) {
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
					goto LOOP_INIT_bmzxaf
				LOOP_INIT_bmzxaf:
					;
					d := 0
					goto LOOP_COND_vulcka
				LOOP_COND_vulcka:
					if d < 1 {
						goto LOOP_BODY_ujvqka
					} else {
						goto LOOP_END_mosjgl
					}
				LOOP_BODY_ujvqka:
					{
						d++
						goto LOOP_COND_vulcka
					}
				LOOP_END_mosjgl:
					{
					}
				}
				query = query.Where(gstr.CaseSnake(c.Column), c.Val)
			case NE:
				{
					goto LOOP_INIT_kurdtk
				LOOP_INIT_kurdtk:
					;
					d := 0
					goto LOOP_COND_ymqnsr
				LOOP_COND_ymqnsr:
					if d < 1 {
						goto LOOP_BODY_kesoen
					} else {
						goto LOOP_END_kpiqiw
					}
				LOOP_BODY_kesoen:
					{
						d++
						goto LOOP_COND_ymqnsr
					}
				LOOP_END_kpiqiw:
					{
					}
				}
				query = query.WhereNot(gstr.CaseSnake(c.Column), c.Val)
			case GT:
				{
					goto LOOP_INIT_dmdytq
				LOOP_INIT_dmdytq:
					;
					d := 0
					goto LOOP_COND_mvcaph
				LOOP_COND_mvcaph:
					if d < 1 {
						goto LOOP_BODY_cposve
					} else {
						goto LOOP_END_erhnao
					}
				LOOP_BODY_cposve:
					{
						d++
						goto LOOP_COND_mvcaph
					}
				LOOP_END_erhnao:
					{
					}
				}
				query = query.WhereGT(gstr.CaseSnake(c.Column), c.Val)
			case GE:
				query = query.WhereGTE(gstr.CaseSnake(c.Column), c.Val)
			case LT:
				{
					goto LOOP_INIT_wtvcst
				LOOP_INIT_wtvcst:
					;
					d := 0
					goto LOOP_COND_qkgzto
				LOOP_COND_qkgzto:
					if d < 1 {
						goto LOOP_BODY_lmkbrz
					} else {
						goto LOOP_END_yfaqco
					}
				LOOP_BODY_lmkbrz:
					{
						d++
						goto LOOP_COND_qkgzto
					}
				LOOP_END_yfaqco:
					{
					}
				}
				query = query.WhereLT(gstr.CaseSnake(c.Column), c.Val)
			case LE:
				query = query.WhereLTE(gstr.CaseSnake(c.Column), c.Val)
			case LIKE:
				{
					goto LOOP_INIT_jlpfzi
				LOOP_INIT_jlpfzi:
					;
					d := 0
					goto LOOP_COND_rvzucw
				LOOP_COND_rvzucw:
					if d < 1 {
						goto LOOP_BODY_zvhbwv
					} else {
						goto LOOP_END_ylbqpt
					}
				LOOP_BODY_zvhbwv:
					{
						query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
						d++
						goto LOOP_COND_rvzucw

					}
				LOOP_END_ylbqpt:
					{
					}
				}
			case NOT_LIKE:
				{
					goto LOOP_INIT_outaix
				LOOP_INIT_outaix:
					;
					d := 0
					goto LOOP_COND_lucqrk
				LOOP_COND_lucqrk:
					if d < 1 {
						goto LOOP_BODY_fpjlfo
					} else {
						goto LOOP_END_ndpnlk
					}
				LOOP_BODY_fpjlfo:
					{
						d++
						goto LOOP_COND_lucqrk
					}
				LOOP_END_ndpnlk:
					{
					}
				}
				query = query.WhereNotLike(gstr.CaseSnake(c.Column), c.Val.(string))
			case LIKE_LEFT:
				{
					goto LOOP_INIT_cmjcyz
				LOOP_INIT_cmjcyz:
					;
					d := 0
					goto LOOP_COND_sxhafy
				LOOP_COND_sxhafy:
					if d < 1 {
						goto LOOP_BODY_rearez
					} else {
						goto LOOP_END_ivsamr
					}
				LOOP_BODY_rearez:
					{
						d++
						goto LOOP_COND_sxhafy
					}
				LOOP_END_ivsamr:
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
					goto LOOP_INIT_hnyqdw
				LOOP_INIT_hnyqdw:
					;
					d := 0
					goto LOOP_COND_jllqlw
				LOOP_COND_jllqlw:
					if d < 1 {
						goto LOOP_BODY_zbeliw
					} else {
						goto LOOP_END_cykojk
					}
				LOOP_BODY_zbeliw:
					{
						d++
						goto LOOP_COND_jllqlw
					}
				LOOP_END_cykojk:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereIn(gstr.CaseSnake(c.Column), c.Val)
				}
			case NOT_IN:
				{
					goto LOOP_INIT_fpfymq
				LOOP_INIT_fpfymq:
					;
					d := 0
					goto LOOP_COND_otqegq
				LOOP_COND_otqegq:
					if d < 1 {
						goto LOOP_BODY_ynmigu
					} else {
						goto LOOP_END_jrrcxq
					}
				LOOP_BODY_ynmigu:
					{
						d++
						goto LOOP_COND_otqegq
					}
				LOOP_END_jrrcxq:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereNotIn(gstr.CaseSnake(c.Column), c.Val)
				}
			case IN_STR:
				if !g.IsEmpty(c.Val) {
					query = query.WhereIn(gstr.CaseSnake(c.Column), gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab")))
				}
			case NOT_IN_STR:
				{
					goto LOOP_INIT_bowhzn
				LOOP_INIT_bowhzn:
					;
					d := 0
					goto LOOP_COND_ezvgkb
				LOOP_COND_ezvgkb:
					if d < 1 {
						goto LOOP_BODY_cgbubp
					} else {
						goto LOOP_END_husudr
					}
				LOOP_BODY_cgbubp:
					{
						d++
						goto LOOP_COND_ezvgkb
					}
				LOOP_END_husudr:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereNotIn(gstr.CaseSnake(c.Column), gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab")))
				}
			case FIND_IN_SET:
				{
					goto LOOP_INIT_qkbzwd
				LOOP_INIT_qkbzwd:
					;
					d := 0
					goto LOOP_COND_vppdsm
				LOOP_COND_vppdsm:
					if d < 1 {
						goto LOOP_BODY_ffswyh
					} else {
						goto LOOP_END_ypkgzx
					}
				LOOP_BODY_ffswyh:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range c.Val.([]uint64) {
								findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_vppdsm

					}
				LOOP_END_ypkgzx:
					{
					}
				}
			case FIND_IN_SET_STR:
				{
					goto LOOP_INIT_rfnohk
				LOOP_INIT_rfnohk:
					;
					d := 0
					goto LOOP_COND_wesheb
				LOOP_COND_wesheb:
					if d < 1 {
						goto LOOP_BODY_viiagd
					} else {
						goto LOOP_END_bhexbe
					}
				LOOP_BODY_viiagd:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range gconv.SliceUint64(gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab"))) {
								findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58c192a8b31144b7778b9071a3892f37fac594bb2c11b96df69"), u, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_wesheb

					}
				LOOP_END_bhexbe:
					{
					}
				}
			case FIND_IN_SET_STR_STR:
				{
					goto LOOP_INIT_sjztft
				LOOP_INIT_sjztft:
					;
					d := 0
					goto LOOP_COND_ykrsms
				LOOP_COND_ykrsms:
					if d < 1 {
						goto LOOP_BODY_haxybt
					} else {
						goto LOOP_END_vbbddf
					}
				LOOP_BODY_haxybt:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range gconv.SliceStr(gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab"))) {
								findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58c193d8b31144b777815b380b3428fd8b9d307af3252056f3c"), u, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_ykrsms

					}
				LOOP_END_vbbddf:
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
		zXXX := int64(1)
		sXXX := float64(6)
		{
			goto LOOP_INIT_dfzygz
		LOOP_INIT_dfzygz:
			;
			iXXX :=

				1
			goto LOOP_COND_exvnbb
		LOOP_COND_exvnbb:
			if iXXX < 15 {
				goto LOOP_BODY_visexe
			} else {
				goto LOOP_END_mcfgyg
			}
		LOOP_BODY_visexe:
			{
				{
					goto LOOP_INIT_stnpvn
				LOOP_INIT_stnpvn:
					;
					jXXX := iXXX
					goto LOOP_COND_ihefrv
				LOOP_COND_ihefrv:
					if jXXX < 15 {
						goto LOOP_BODY_uwxdmm
					} else {
						goto LOOP_END_akaffq
					}
				LOOP_BODY_uwxdmm:
					{
						{
							goto LOOP_INIT_wsjowg
						LOOP_INIT_wsjowg:
							;
							zXXX := jXXX
							goto LOOP_COND_yajcfy
						LOOP_COND_yajcfy:
							if zXXX < 15 {
								goto LOOP_BODY_zspagn
							} else {
								goto LOOP_END_usnuqw
							}
						LOOP_BODY_zspagn:

							{
								sXXX = (float64(iXXX+jXXX) * float64(zXXX)) / float64(iXXX)
								zXXX++
								goto LOOP_COND_yajcfy

							}
						LOOP_END_usnuqw:
							{
							}
						}
						jXXX++
						goto LOOP_COND_ihefrv

					}
				LOOP_END_akaffq:
					{
					}
				}
				iXXX++
				goto LOOP_COND_exvnbb

			}
		LOOP_END_mcfgyg:
			{
			}
		}
		if sXXX ==
			float64(zXXX) {
		}
	})()

	if !g.IsEmpty(sidx) {
		{
			goto LOOP_INIT_vsfxfx
		LOOP_INIT_vsfxfx:
			;
			d := 0
			goto LOOP_COND_ikydca
		LOOP_COND_ikydca:
			if d < 1 {
				goto LOOP_BODY_wwcmft
			} else {
				goto LOOP_END_ymixoa
			}
		LOOP_BODY_wwcmft:
			{
				if gstr.Equal(sort, ORDER_BY_DESC) {
					{
						goto LOOP_INIT_pvcljc
					LOOP_INIT_pvcljc:
						;
						d := 0
						goto LOOP_COND_vnxyyd
					LOOP_COND_vnxyyd:
						if d < 1 {
							goto LOOP_BODY_oklkve
						} else {
							goto LOOP_END_sjvurf
						}
					LOOP_BODY_oklkve:
						{
							query = query.OrderDesc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_vnxyyd

						}
					LOOP_END_sjvurf:
						{
						}
					}
				} else {
					{
						goto LOOP_INIT_zqnvlm
					LOOP_INIT_zqnvlm:
						;
						d := 0
						goto LOOP_COND_zqkvws
					LOOP_COND_zqkvws:
						if d < 1 {
							goto LOOP_BODY_bszfet
						} else {
							goto LOOP_END_sazywh
						}
					LOOP_BODY_bszfet:
						{
							query = query.OrderAsc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_zqkvws

						}
					LOOP_END_sazywh:
						{
						}
					}
				}
				d++
				goto LOOP_COND_ikydca

			}
		LOOP_END_ymixoa:
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
