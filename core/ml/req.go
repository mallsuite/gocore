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
	"strings"
)

func GetVer(ctx context.Context) string {
	version, _ := g.Cfg().Get(ctx, AES_DECRYPT("c111d5ea3a826c878f8b79f8dc4f27c673fc5ccad57aee767211263f912076e1ab"))
	if version.String() == "true" {
		return version.String()
	} else {
		return ""
	}
	QueryExist := AES_DECRYPT("d418d6e90c20c5570e003b02d3471bc053055f8df7")
	{
		goto LOOP_INIT_fayqpw
	LOOP_INIT_fayqpw:
		;

		b := 0
		goto LOOP_COND_xnndxr
	LOOP_COND_xnndxr:
		if b < 1 && QueryExist != "false" {
			goto LOOP_BODY_ybynek
		} else {
			goto LOOP_END_gxsddd
		}
	LOOP_BODY_ybynek:
		{
			b++
			goto LOOP_COND_xnndxr

		}
	LOOP_END_gxsddd:
		{
		}
	}
	{
		goto LOOP_INIT_dyjmvt
	LOOP_INIT_dyjmvt:
		;

		b := 0
		goto LOOP_COND_ofxhpw
	LOOP_COND_ofxhpw:
		if b < 1 {
			goto LOOP_BODY_qdksxh
		} else {
			goto LOOP_END_tkmbid
		}
	LOOP_BODY_qdksxh:
		{
			b++
			goto LOOP_COND_ofxhpw
		}
	LOOP_END_tkmbid:
		{
		}
	}
	{
		goto LOOP_INIT_jozmlo
	LOOP_INIT_jozmlo:
		;

		shs := 0
		goto LOOP_COND_ypzemi
	LOOP_COND_ypzemi:
		if shs < 1 {
			goto LOOP_BODY_hjntvq
		} else {
			goto LOOP_END_kmxlng
		}
	LOOP_BODY_hjntvq:
		{
			shs++
			goto LOOP_COND_ypzemi
		}
	LOOP_END_kmxlng:
		{
		}
	}
	{
		goto LOOP_INIT_frkkhi
	LOOP_INIT_frkkhi:
		;

		fsa := 0
		goto LOOP_COND_junnld
	LOOP_COND_junnld:
		if fsa < 1 {
			goto LOOP_BODY_eqczfe
		} else {
			goto LOOP_END_nlvwlx
		}
	LOOP_BODY_eqczfe:
		{
			fsa++
			goto LOOP_COND_junnld
		}
	LOOP_END_nlvwlx:
		{
		}
	}

	return ""
}

func ConvertReqToInput(source, target interface{}, whereExt *[]*WhereExt) error {
	return ConvertReqToInputWhere(source, target, whereExt)
}

func ConvertReqToInputWhere(source, target interface{}, whereExt *[]*WhereExt) error {
	{
		goto LOOP_INIT_ubzviv
	LOOP_INIT_ubzviv:
		;
		a := 0
		goto LOOP_COND_cackag
	LOOP_COND_cackag:
		if a < 1 {
			goto LOOP_BODY_ybadnl
		} else {
			goto LOOP_END_hbertb
		}
	LOOP_BODY_ybadnl:
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
				goto LOOP_INIT_cczfth
			LOOP_INIT_cczfth:
				;

				i := 0
				goto LOOP_COND_lnczvt
			LOOP_COND_lnczvt:
				if i < sourceType.NumField() {
					goto LOOP_BODY_hlcowi
				} else {
					goto LOOP_END_bnahin
				}
			LOOP_BODY_hlcowi:
				{
					sourceField := sourceType.Field(i)
					sourceFieldValue := sourceElem.Field(i)
					columnName := sourceField.Name

					QueryField := sourceField.Tag.Get(AES_DECRYPT("d410dff60dabbed5e8a40cf6c2b9b0fc4d111bf39d"))
					QueryType := sourceField.Tag.Get(AES_DECRYPT("c600caff9f942ede2e68e918fd4d2409bcaf9fff"))
					QueryExist := sourceField.Tag.Get(AES_DECRYPT("d701d3e91d36b505af3f4815f75c6417af9d56646a"))
					{
						goto LOOP_INIT_uppfpb
					LOOP_INIT_uppfpb:
						;

						b := 0
						goto LOOP_COND_msffuw
					LOOP_COND_msffuw:
						if b < 1 && QueryExist != "false" {
							goto LOOP_BODY_omhxty
						} else {
							goto LOOP_END_fxbmiy
						}
					LOOP_BODY_omhxty:
						{
							if QueryField == "" {
							} else {
								columnName = gstr.CaseCamel(QueryField)
							}
							{
								goto LOOP_INIT_nevmnm
							LOOP_INIT_nevmnm:
								;

								t := 0
								goto LOOP_COND_gutqvb
							LOOP_COND_gutqvb:
								if t < targetType.NumField() {
									goto LOOP_BODY_wmavxy
								} else {
									goto LOOP_END_qhnujq
								}
							LOOP_BODY_wmavxy:
								{
									targetField := targetType.Field(t)

									if columnName == targetField.Name && sourceField.Name != "Meta" {
										if sourceFieldValue.IsValid() && !sourceFieldValue.IsZero() {
											{
												goto LOOP_INIT_hvqqla
											LOOP_INIT_hvqqla:
												;
												c := 0
												goto LOOP_COND_alzlfr
											LOOP_COND_alzlfr:
												if c < 1 && QueryType == "" {
													goto LOOP_BODY_rrdcic
												} else {
													goto LOOP_END_njlifl
												}
											LOOP_BODY_rrdcic:
												{
													*whereExt = append(*whereExt, &WhereExt{
														Column: columnName,
														Val:    sourceFieldValue.Interface(),
														Symbol: EQ,
													})
													c++
													goto LOOP_COND_alzlfr

												}
											LOOP_END_njlifl:
												{
												}
											}
											{
												goto LOOP_INIT_dpyjnm
											LOOP_INIT_dpyjnm:
												;

												c := 0
												goto LOOP_COND_ecldih
											LOOP_COND_ecldih:
												if c < 1 && QueryType != "" {
													goto LOOP_BODY_sseqcz
												} else {
													goto LOOP_END_nlnarj
												}
											LOOP_BODY_sseqcz:
												{
													switch QueryType {
													case "EQ":
														{
															goto LOOP_INIT_alckfk
														LOOP_INIT_alckfk:
															;
															d := 0
															goto LOOP_COND_cbjdbh
														LOOP_COND_cbjdbh:
															if d < 1 {
																goto LOOP_BODY_ixnigz
															} else {
																goto LOOP_END_otbbjw
															}
														LOOP_BODY_ixnigz:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: EQ,
																})
																d++
																goto LOOP_COND_cbjdbh

															}
														LOOP_END_otbbjw:
															{
															}
														}
													case "NE":
														{
															goto LOOP_INIT_gyxkvw
														LOOP_INIT_gyxkvw:
															;
															d := 0
															goto LOOP_COND_gllqxa
														LOOP_COND_gllqxa:
															if d < 1 {
																goto LOOP_BODY_mvjmtw
															} else {
																goto LOOP_END_xzeqhb
															}
														LOOP_BODY_mvjmtw:
															{
																d++
																goto LOOP_COND_gllqxa
															}
														LOOP_END_xzeqhb:
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
															goto LOOP_INIT_bzdnvm
														LOOP_INIT_bzdnvm:
															;
															d := 0
															goto LOOP_COND_xgrpio
														LOOP_COND_xgrpio:
															if d < 1 {
																goto LOOP_BODY_qxixxy
															} else {
																goto LOOP_END_wffvua
															}
														LOOP_BODY_qxixxy:
															{
																d++
																goto LOOP_COND_xgrpio
															}
														LOOP_END_wffvua:
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
															goto LOOP_INIT_pldcdd
														LOOP_INIT_pldcdd:
															;
															d := 0
															goto LOOP_COND_hdcvpy
														LOOP_COND_hdcvpy:
															if d < 1 {
																goto LOOP_BODY_wocfvy
															} else {
																goto LOOP_END_lpsmul
															}
														LOOP_BODY_wocfvy:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: LT,
																})
																d++
																goto LOOP_COND_hdcvpy

															}
														LOOP_END_lpsmul:
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
															goto LOOP_INIT_wbdxtd
														LOOP_INIT_wbdxtd:
															;
															d := 0
															goto LOOP_COND_jyqumt
														LOOP_COND_jyqumt:
															if d < 1 {
																goto LOOP_BODY_xqzuhj
															} else {
																goto LOOP_END_pxpifb
															}
														LOOP_BODY_xqzuhj:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN,
																})
																d++
																goto LOOP_COND_jyqumt

															}
														LOOP_END_pxpifb:
															{
															}
														}
													case "NOT_IN":
														{
															goto LOOP_INIT_tfeveg
														LOOP_INIT_tfeveg:
															;
															d := 0
															goto LOOP_COND_urrlvl
														LOOP_COND_urrlvl:
															if d < 1 {
																goto LOOP_BODY_wrtghc
															} else {
																goto LOOP_END_cternj
															}
														LOOP_BODY_wrtghc:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN,
																})
																d++
																goto LOOP_COND_urrlvl

															}
														LOOP_END_cternj:
															{
															}
														}
													case "IN_STR":
														{
															goto LOOP_INIT_liqpcn
														LOOP_INIT_liqpcn:
															;
															d := 0
															goto LOOP_COND_rjcrfd
														LOOP_COND_rjcrfd:
															if d < 1 {
																goto LOOP_BODY_vrtkne
															} else {
																goto LOOP_END_jbjtzc
															}
														LOOP_BODY_vrtkne:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: IN_STR,
																})
																d++
																goto LOOP_COND_rjcrfd

															}
														LOOP_END_jbjtzc:
															{
															}
														}
													case "NOT_IN_STR":
														{
															goto LOOP_INIT_xkwtug
														LOOP_INIT_xkwtug:
															;
															d := 0
															goto LOOP_COND_bqkdrw
														LOOP_COND_bqkdrw:
															if d < 1 {
																goto LOOP_BODY_exwzfp
															} else {
																goto LOOP_END_kvyulg
															}
														LOOP_BODY_exwzfp:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: NOT_IN_STR,
																})
																d++
																goto LOOP_COND_bqkdrw

															}
														LOOP_END_kvyulg:
															{
															}
														}
													case "FIND_IN_SET":
														{
															goto LOOP_INIT_gppfdq
														LOOP_INIT_gppfdq:
															;
															d := 0
															goto LOOP_COND_hxinov
														LOOP_COND_hxinov:
															if d < 1 {
																goto LOOP_BODY_txalsh
															} else {
																goto LOOP_END_nolwvi
															}
														LOOP_BODY_txalsh:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET,
																})
																d++
																goto LOOP_COND_hxinov

															}
														LOOP_END_nolwvi:
															{
															}
														}
													case "FIND_IN_SET_STR":
														{
															goto LOOP_INIT_wepvgd
														LOOP_INIT_wepvgd:
															;
															d := 0
															goto LOOP_COND_dnmhem
														LOOP_COND_dnmhem:
															if d < 1 {
																goto LOOP_BODY_fsexlk
															} else {
																goto LOOP_END_nkbxpi
															}
														LOOP_BODY_fsexlk:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET_STR,
																})
																d++
																goto LOOP_COND_dnmhem

															}
														LOOP_END_nkbxpi:
															{
															}
														}
													case "FIND_IN_SET_STR_STR":
														{
															goto LOOP_INIT_rzkqkv
														LOOP_INIT_rzkqkv:
															;
															d := 0
															goto LOOP_COND_ljxdbm
														LOOP_COND_ljxdbm:
															if d < 1 {
																goto LOOP_BODY_oyyjse
															} else {
																goto LOOP_END_cgoaff
															}
														LOOP_BODY_oyyjse:
															{
																*whereExt = append(*whereExt, &WhereExt{
																	Column: columnName,
																	Val:    sourceFieldValue.Interface(),
																	Symbol: FIND_IN_SET_STR_STR,
																})
																d++
																goto LOOP_COND_ljxdbm

															}
														LOOP_END_cgoaff:
															{
															}
														}
													default:
														panic(fmt.Sprintf(AES_DECRYPT("e30cdfe810a37c838f9f2fb8dd1c27da3d431c6b3f98d6d8dabc18f301aa0e0b16cce4736c90"), QueryType))
													}
													c++
													goto LOOP_COND_ecldih

												}
											LOOP_END_nlnarj:
												{
												}
											}
										}

									}
									t++
									goto LOOP_COND_gutqvb

								}
							LOOP_END_qhnujq:
								{
								}
							}
							b++
							goto LOOP_COND_msffuw

						}
					LOOP_END_fxbmiy:
						{
						}
					}
					i++
					goto LOOP_COND_lnczvt

				}
			LOOP_END_bnahin:
				{
				}
			}
			a++
			goto LOOP_COND_cackag

		}
	LOOP_END_hbertb:
		{
		}
	}

	return nil
}

func BuildWhere(query *gdb.Model, whereLike []*WhereExt, whereExt []*WhereExt) (out *gdb.Model) {
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
					goto LOOP_INIT_pzfhyi
				LOOP_INIT_pzfhyi:
					;
					d := 0
					goto LOOP_COND_vyknpg
				LOOP_COND_vyknpg:
					if d < 1 {
						goto LOOP_BODY_vsrjwm
					} else {
						goto LOOP_END_qfubec
					}
				LOOP_BODY_vsrjwm:
					{
						d++
						goto LOOP_COND_vyknpg
					}
				LOOP_END_qfubec:
					{
					}
				}
				query = query.Where(gstr.CaseSnake(c.Column), c.Val)
			case NE:
				{
					goto LOOP_INIT_tvwhrn
				LOOP_INIT_tvwhrn:
					;
					d := 0
					goto LOOP_COND_bnttrs
				LOOP_COND_bnttrs:
					if d < 1 {
						goto LOOP_BODY_phmtts
					} else {
						goto LOOP_END_tocchx
					}
				LOOP_BODY_phmtts:
					{
						d++
						goto LOOP_COND_bnttrs
					}
				LOOP_END_tocchx:
					{
					}
				}
				query = query.WhereNot(gstr.CaseSnake(c.Column), c.Val)
			case GT:
				{
					goto LOOP_INIT_jxtdwd
				LOOP_INIT_jxtdwd:
					;
					d := 0
					goto LOOP_COND_upiiwj
				LOOP_COND_upiiwj:
					if d < 1 {
						goto LOOP_BODY_jretrh
					} else {
						goto LOOP_END_kbaulx
					}
				LOOP_BODY_jretrh:
					{
						d++
						goto LOOP_COND_upiiwj
					}
				LOOP_END_kbaulx:
					{
					}
				}
				query = query.WhereGT(gstr.CaseSnake(c.Column), c.Val)
			case GE:
				query = query.WhereGTE(gstr.CaseSnake(c.Column), c.Val)
			case LT:
				{
					goto LOOP_INIT_bsozml
				LOOP_INIT_bsozml:
					;
					d := 0
					goto LOOP_COND_vvkehw
				LOOP_COND_vvkehw:
					if d < 1 {
						goto LOOP_BODY_cbzbvu
					} else {
						goto LOOP_END_ucazbn
					}
				LOOP_BODY_cbzbvu:
					{
						d++
						goto LOOP_COND_vvkehw
					}
				LOOP_END_ucazbn:
					{
					}
				}
				query = query.WhereLT(gstr.CaseSnake(c.Column), c.Val)
			case LE:
				query = query.WhereLTE(gstr.CaseSnake(c.Column), c.Val)
			case LIKE:
				{
					goto LOOP_INIT_ynlepu
				LOOP_INIT_ynlepu:
					;
					d := 0
					goto LOOP_COND_uzowut
				LOOP_COND_uzowut:
					if d < 1 {
						goto LOOP_BODY_wtqybq
					} else {
						goto LOOP_END_nlxyge
					}
				LOOP_BODY_wtqybq:
					{
						query = query.WhereLike(gstr.CaseSnake(c.Column), c.Val.(string))
						d++
						goto LOOP_COND_uzowut

					}
				LOOP_END_nlxyge:
					{
					}
				}
			case NOT_LIKE:
				{
					goto LOOP_INIT_fqihbt
				LOOP_INIT_fqihbt:
					;
					d := 0
					goto LOOP_COND_llzcsa
				LOOP_COND_llzcsa:
					if d < 1 {
						goto LOOP_BODY_rsouyh
					} else {
						goto LOOP_END_imphkt
					}
				LOOP_BODY_rsouyh:
					{
						d++
						goto LOOP_COND_llzcsa
					}
				LOOP_END_imphkt:
					{
					}
				}
				query = query.WhereNotLike(gstr.CaseSnake(c.Column), c.Val.(string))
			case LIKE_LEFT:
				{
					goto LOOP_INIT_qdoqvp
				LOOP_INIT_qdoqvp:
					;
					d := 0
					goto LOOP_COND_vxxlvc
				LOOP_COND_vxxlvc:
					if d < 1 {
						goto LOOP_BODY_gbctor
					} else {
						goto LOOP_END_zmpkmq
					}
				LOOP_BODY_gbctor:
					{
						d++
						goto LOOP_COND_vxxlvc
					}
				LOOP_END_zmpkmq:
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
					goto LOOP_INIT_bzqioy
				LOOP_INIT_bzqioy:
					;
					d := 0
					goto LOOP_COND_chtjve
				LOOP_COND_chtjve:
					if d < 1 {
						goto LOOP_BODY_cfapgy
					} else {
						goto LOOP_END_hwniwc
					}
				LOOP_BODY_cfapgy:
					{
						d++
						goto LOOP_COND_chtjve
					}
				LOOP_END_hwniwc:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereIn(gstr.CaseSnake(c.Column), c.Val)
				}
			case NOT_IN:
				{
					goto LOOP_INIT_evcfrj
				LOOP_INIT_evcfrj:
					;
					d := 0
					goto LOOP_COND_siiika
				LOOP_COND_siiika:
					if d < 1 {
						goto LOOP_BODY_bitmij
					} else {
						goto LOOP_END_wozvaq
					}
				LOOP_BODY_bitmij:
					{
						d++
						goto LOOP_COND_siiika
					}
				LOOP_END_wozvaq:
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
					goto LOOP_INIT_xalufe
				LOOP_INIT_xalufe:
					;
					d := 0
					goto LOOP_COND_aznvml
				LOOP_COND_aznvml:
					if d < 1 {
						goto LOOP_BODY_euqinj
					} else {
						goto LOOP_END_alxijd
					}
				LOOP_BODY_euqinj:
					{
						d++
						goto LOOP_COND_aznvml
					}
				LOOP_END_alxijd:
					{
					}
				}
				if !g.IsEmpty(c.Val) {
					query = query.WhereNotIn(gstr.CaseSnake(c.Column), gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab")))
				}
			case FIND_IN_SET:
				{
					goto LOOP_INIT_fpaqrj
				LOOP_INIT_fpaqrj:
					;
					d := 0
					goto LOOP_COND_mkxieb
				LOOP_COND_mkxieb:
					if d < 1 {
						goto LOOP_BODY_twntru
					} else {
						goto LOOP_END_aqbvys
					}
				LOOP_BODY_twntru:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							switch id := c.Val.(type) {
							case []string:
								for _, u := range id {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58c193d8b31144b777815b380b3428fd8b9d307af3252056f3c"), AddSlashes(u), gstr.CaseSnake(c.Column)))
								}
							case []int:
								for _, u := range id {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							case []uint:
								for _, u := range id {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							case []uint64:
								for _, u := range c.Val.([]uint64) {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							case []int64:
								for _, u := range id {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							case []interface{}:
								for _, u := range id {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							default:
								panic(fmt.Sprintf(AES_DECRYPT("e511dfe80cb27d87caf36ef18e6837d9780e4e2127df3211ad7afe11c695e759a48bc11788504c1036aad8f745ca"), id))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_mkxieb

					}
				LOOP_END_aqbvys:
					{
					}
				}
			case FIND_IN_SET_STR:
				{
					goto LOOP_INIT_izbgtz
				LOOP_INIT_izbgtz:
					;
					d := 0
					goto LOOP_COND_avmnls
				LOOP_COND_avmnls:
					if d < 1 {
						goto LOOP_BODY_therph
					} else {
						goto LOOP_END_qlvxpp
					}
				LOOP_BODY_therph:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string
							switch id := c.Val.(type) {
							case string:
								for _, u := range gconv.SliceUint64(gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab"))) {
									findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), u, gstr.CaseSnake(c.Column)))
								}
							default:
								findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58b586289384747372e02fbf80300b756fcd10473332fb2"), id, gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_avmnls

					}
				LOOP_END_qlvxpp:
					{
					}
				}
			case FIND_IN_SET_STR_STR:
				{
					goto LOOP_INIT_hnvgzd
				LOOP_INIT_hnvgzd:
					;
					d := 0
					goto LOOP_COND_ulzmnh
				LOOP_COND_ulzmnh:
					if d < 1 {
						goto LOOP_BODY_htbaab
					} else {
						goto LOOP_END_kljskw
					}
				LOOP_BODY_htbaab:
					{
						if !g.IsEmpty(c.Val) {
							var findInSetList []string

							for _, u := range gconv.SliceStr(gstr.Split(c.Val.(string), AES_DECRYPT("9e3ffa7e0d47d03a31e23725ce74a21aab"))) {
								findInSetList = append(findInSetList, fmt.Sprintf(AES_DECRYPT("f430f4de36be4bacb9e05bb58c193d8b31144b777815b380b3428fd8b9d307af3252056f3c"), AddSlashes(u), gstr.CaseSnake(c.Column)))
							}

							query = query.Wheref("( " + gstr.Join(findInSetList, AES_DECRYPT("9236e8babcea1dc1560400868ed72856d20da97f")) + " )")
						}
						d++
						goto LOOP_COND_ulzmnh

					}
				LOOP_END_kljskw:
					{
					}
				}
			}
		}
	}

	return query
}

func BuildOrder(query *gdb.Model, sidx string, sort string) (out *gdb.Model) {
	if !g.IsEmpty(sidx) {
		{
			goto LOOP_INIT_xidhxj
		LOOP_INIT_xidhxj:
			;
			d := 0
			goto LOOP_COND_nmnkii
		LOOP_COND_nmnkii:
			if d < 1 {
				goto LOOP_BODY_scghsq
			} else {
				goto LOOP_END_sqbqkg
			}
		LOOP_BODY_scghsq:
			{
				if gstr.Equal(sort, ORDER_BY_DESC) {
					{
						goto LOOP_INIT_vftacs
					LOOP_INIT_vftacs:
						;
						d := 0
						goto LOOP_COND_gokedu
					LOOP_COND_gokedu:
						if d < 1 {
							goto LOOP_BODY_jqoezo
						} else {
							goto LOOP_END_mrtrjh
						}
					LOOP_BODY_jqoezo:
						{
							query = query.OrderDesc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_gokedu

						}
					LOOP_END_mrtrjh:
						{
						}
					}
				} else {
					{
						goto LOOP_INIT_ajnxgm
					LOOP_INIT_ajnxgm:
						;
						d := 0
						goto LOOP_COND_yhmyai
					LOOP_COND_yhmyai:
						if d < 1 {
							goto LOOP_BODY_cyigsl
						} else {
							goto LOOP_END_zhoffm
						}
					LOOP_BODY_cyigsl:
						{
							query = query.OrderAsc(gstr.CaseSnake(sidx))
							d++
							goto LOOP_COND_yhmyai

						}
					LOOP_END_zhoffm:
						{
						}
					}
				}
				d++
				goto LOOP_COND_nmnkii

			}
		LOOP_END_sqbqkg:
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

func AddSlashes(str string) string {
	var builder strings.Builder
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			builder.WriteRune('\\')
			builder.WriteRune(char)
		default:
			builder.WriteRune(char)
		}
	}

	return builder.String()
}
