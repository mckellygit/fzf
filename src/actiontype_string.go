// Code generated by "stringer -type=actionType"; DO NOT EDIT.

package fzf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[actIgnore-0]
	_ = x[actStart-1]
	_ = x[actClick-2]
	_ = x[actInvalid-3]
	_ = x[actChar-4]
	_ = x[actMouse-5]
	_ = x[actBeginningOfLine-6]
	_ = x[actAbort-7]
	_ = x[actAccept-8]
	_ = x[actAcceptNonEmpty-9]
	_ = x[actAcceptOrPrintQuery-10]
	_ = x[actBackwardChar-11]
	_ = x[actBackwardDeleteChar-12]
	_ = x[actBackwardDeleteCharEof-13]
	_ = x[actBackwardWord-14]
	_ = x[actCancel-15]
	_ = x[actChangeBorderLabel-16]
	_ = x[actChangeHeader-17]
	_ = x[actChangeMulti-18]
	_ = x[actChangePreviewLabel-19]
	_ = x[actChangePrompt-20]
	_ = x[actChangeQuery-21]
	_ = x[actClearScreen-22]
	_ = x[actClearQuery-23]
	_ = x[actClearSelection-24]
	_ = x[actClose-25]
	_ = x[actDeleteChar-26]
	_ = x[actDeleteCharEof-27]
	_ = x[actEndOfLine-28]
	_ = x[actFatal-29]
	_ = x[actForwardChar-30]
	_ = x[actForwardWord-31]
	_ = x[actKillLine-32]
	_ = x[actKillWord-33]
	_ = x[actUnixLineDiscard-34]
	_ = x[actUnixWordRubout-35]
	_ = x[actYank-36]
	_ = x[actBackwardKillWord-37]
	_ = x[actSelectAll-38]
	_ = x[actDeselectAll-39]
	_ = x[actToggle-40]
	_ = x[actToggleSearch-41]
	_ = x[actToggleAll-42]
	_ = x[actToggleDown-43]
	_ = x[actToggleUp-44]
	_ = x[actToggleIn-45]
	_ = x[actToggleOut-46]
	_ = x[actToggleTrack-47]
	_ = x[actToggleTrackCurrent-48]
	_ = x[actToggleHeader-49]
	_ = x[actToggleWrap-50]
	_ = x[actTrackCurrent-51]
	_ = x[actUntrackCurrent-52]
	_ = x[actDown-53]
	_ = x[actUp-54]
	_ = x[actPageUp-55]
	_ = x[actPageDown-56]
	_ = x[actPosition-57]
	_ = x[actHalfPageUp-58]
	_ = x[actHalfPageDown-59]
	_ = x[actOffsetUp-60]
	_ = x[actOffsetDown-61]
	_ = x[actOffsetMiddle-62]
	_ = x[actJump-63]
	_ = x[actJumpAccept-64]
	_ = x[actPrintQuery-65]
	_ = x[actRefreshPreview-66]
	_ = x[actReplaceQuery-67]
	_ = x[actToggleSort-68]
	_ = x[actShowPreview-69]
	_ = x[actHidePreview-70]
	_ = x[actTogglePreview-71]
	_ = x[actTogglePreviewWrap-72]
	_ = x[actTransform-73]
	_ = x[actTransformBorderLabel-74]
	_ = x[actTransformHeader-75]
	_ = x[actTransformPreviewLabel-76]
	_ = x[actTransformPrompt-77]
	_ = x[actTransformQuery-78]
	_ = x[actPreview-79]
	_ = x[actChangePreview-80]
	_ = x[actChangePreviewWindow-81]
	_ = x[actPreviewTop-82]
	_ = x[actPreviewBottom-83]
	_ = x[actPreviewUp-84]
	_ = x[actPreviewDown-85]
	_ = x[actPreviewPageUp-86]
	_ = x[actPreviewPageDown-87]
	_ = x[actPreviewHalfPageUp-88]
	_ = x[actPreviewHalfPageDown-89]
	_ = x[actPrevHistory-90]
	_ = x[actPrevSelected-91]
	_ = x[actPrint-92]
	_ = x[actPut-93]
	_ = x[actNextHistory-94]
	_ = x[actNextSelected-95]
	_ = x[actExecute-96]
	_ = x[actExecuteSilent-97]
	_ = x[actExecuteMulti-98]
	_ = x[actSigStop-99]
	_ = x[actFirst-100]
	_ = x[actLast-101]
	_ = x[actReload-102]
	_ = x[actReloadSync-103]
	_ = x[actDisableSearch-104]
	_ = x[actEnableSearch-105]
	_ = x[actSelect-106]
	_ = x[actDeselect-107]
	_ = x[actUnbind-108]
	_ = x[actRebind-109]
	_ = x[actBecome-110]
	_ = x[actShowHeader-111]
	_ = x[actHideHeader-112]
}

const _actionType_name = "actIgnoreactStartactClickactInvalidactCharactMouseactBeginningOfLineactAbortactAcceptactAcceptNonEmptyactAcceptOrPrintQueryactBackwardCharactBackwardDeleteCharactBackwardDeleteCharEofactBackwardWordactCancelactChangeBorderLabelactChangeHeaderactChangeMultiactChangePreviewLabelactChangePromptactChangeQueryactClearScreenactClearQueryactClearSelectionactCloseactDeleteCharactDeleteCharEofactEndOfLineactFatalactForwardCharactForwardWordactKillLineactKillWordactUnixLineDiscardactUnixWordRuboutactYankactBackwardKillWordactSelectAllactDeselectAllactToggleactToggleSearchactToggleAllactToggleDownactToggleUpactToggleInactToggleOutactToggleTrackactToggleTrackCurrentactToggleHeaderactToggleWrapactTrackCurrentactUntrackCurrentactDownactUpactPageUpactPageDownactPositionactHalfPageUpactHalfPageDownactOffsetUpactOffsetDownactOffsetMiddleactJumpactJumpAcceptactPrintQueryactRefreshPreviewactReplaceQueryactToggleSortactShowPreviewactHidePreviewactTogglePreviewactTogglePreviewWrapactTransformactTransformBorderLabelactTransformHeaderactTransformPreviewLabelactTransformPromptactTransformQueryactPreviewactChangePreviewactChangePreviewWindowactPreviewTopactPreviewBottomactPreviewUpactPreviewDownactPreviewPageUpactPreviewPageDownactPreviewHalfPageUpactPreviewHalfPageDownactPrevHistoryactPrevSelectedactPrintactPutactNextHistoryactNextSelectedactExecuteactExecuteSilentactExecuteMultiactSigStopactFirstactLastactReloadactReloadSyncactDisableSearchactEnableSearchactSelectactDeselectactUnbindactRebindactBecomeactShowHeaderactHideHeader"

var _actionType_index = [...]uint16{0, 9, 17, 25, 35, 42, 50, 68, 76, 85, 102, 123, 138, 159, 183, 198, 207, 227, 242, 256, 277, 292, 306, 320, 333, 350, 358, 371, 387, 399, 407, 421, 435, 446, 457, 475, 492, 499, 518, 530, 544, 553, 568, 580, 593, 604, 615, 627, 641, 662, 677, 690, 705, 722, 729, 734, 743, 754, 765, 778, 793, 804, 817, 832, 839, 852, 865, 882, 897, 910, 924, 938, 954, 974, 986, 1009, 1027, 1051, 1069, 1086, 1096, 1112, 1134, 1147, 1163, 1175, 1189, 1205, 1223, 1243, 1265, 1279, 1294, 1302, 1308, 1322, 1337, 1347, 1363, 1378, 1388, 1396, 1403, 1412, 1425, 1441, 1456, 1465, 1476, 1485, 1494, 1503, 1516, 1529}

func (i actionType) String() string {
	if i < 0 || i >= actionType(len(_actionType_index)-1) {
		return "actionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _actionType_name[_actionType_index[i]:_actionType_index[i+1]]
}
