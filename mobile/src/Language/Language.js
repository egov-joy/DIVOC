import React, {useState} from "react";
import {BaseCard} from "../Base/Base";
import "./Language.scss"
import {Button} from "react-bootstrap";
import * as PropTypes from "prop-types";
import {getMessageComponent, getNumberComponent, LANGUAGE_KEYS, useLocale} from "../lang/LocaleContext";

const languageSupports = [
    {
        languageName: "English",
        languageCode: "en"
    },
    {
        languageName: "Hindi",
        languageCode: "hi"
    }/*,
    {
        languageName: "Arabic",
        languageCode: "ar"
    }*/
]


function LanguageSelection(props) {
    const locale = useLocale()

    return <div className={"language-container"}>
        <h1 className={"title"}>{getMessageComponent(LANGUAGE_KEYS.SELECT_LANGUAGE)}</h1>
        {languageSupports.map((item, index) => {
            return <Button onClick={() => locale.selectLanguage(item.languageCode)}>{item.languageName}</Button>
        })}
    </div>;
}

LanguageSelection.propTypes = {
    onClick: PropTypes.func,
    onClick1: PropTypes.func
};

export function SelectLanguage(props) {
    return (
        <BaseCard>
            <LanguageSelection/>
        </BaseCard>
    );

}
