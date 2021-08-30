<#import "template.ftl" as layout>
    <@layout.registrationLayout displayInfo=social.displayInfo displayWide=(realm.password && social.providers??);
        section>
        <#if section="header">

            <#elseif section="form">
                <div id="kc-form" <#if realm.password && social.providers??>class="${properties.kcContentWrapperClass!}"
        </#if>>
        <div id="kc-form-wrapper" <#if realm.password && social.providers??>
            class="${properties.kcFormSocialAccountContentClass!} ${properties.kcFormSocialAccountClass!}"</#if>>
            <#if realm.password>
                <form id="kc-form-login" onsubmit="login.disabled = true; return true;" action="${url.loginAction}"
                    method="post">

                    <div>
                        <label class="inputlabel" for="username">
                            <#if !realm.loginWithEmailAllowed>${msg("username")}<#elseif
                                    !realm.registrationEmailAsUsername>${msg("usernameOrEmail")}<#else>${msg("email")}
                            </#if>
                        </label> <br />

                        <#if usernameEditDisabled??>
                            <input tabindex="1" class="inputs" id="username" name="username"
                                value="${(login.username!'')}" type="text" disabled />
                            <#else>
                                <input tabindex="1" class="inputs" id="username" name="username"
                                    value="${(login.username!'')}" type="text" autofocus autocomplete="off" />
                        </#if>
                    </div>

                    <div>
                        <label class="inputlabel" for="password">${msg("password")}</label> <br />
                        <input tabindex="2" class="inputs" id="password" class="${properties.kcInputClass!}"
                            name="password" type="password" autocomplete="off" />
                    </div>

                    <div class="${properties.kcFormGroupClass!} ${properties.kcFormSettingClass!}">
                        <div id="kc-form-options">
                            <#if realm.rememberMe && !usernameEditDisabled??>

                                <div class="mdc-form-field">
                                    <div class="mdc-checkbox">
                                        <input tabindex="0" id="rememberMe" name="rememberMe" type="checkbox"
                                            class="mdc-checkbox__native-control" <#if login.rememberMe??>checked
                            </#if>
                            />
                            <div class="mdc-checkbox__background">
                                <svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
                                    <path class="mdc-checkbox__checkmark-path" fill="none"
                                        d="M1.73,12.91 8.1,19.28 22.79,4.59" />
                                </svg>
                                <div class="mdc-checkbox__mixedmark"></div>
                            </div>
                            <div class="mdc-checkbox__ripple"></div>
                        </div>
                        <label for="rememberMe">${msg("rememberMe")}</label>
                    </div>

            </#if>
        </div>
        <div class="${properties.kcFormOptionsWrapperClass!}">
            <#if realm.resetPasswordAllowed>
                <span><a tabindex="0" href="${url.loginResetCredentialsUrl}">${msg("doForgotPassword")}</a></span>
            </#if>
        </div>
        </div>

        <#-- <div id="kc-form-buttons" class="${properties.kcFormGroupClass!}">
            <input type="hidden" id="id-hidden-input" name="credentialId" <#if
                auth.selectedCredential?has_content>value="${auth.selectedCredential}"</#if>/>
            <input tabindex="4"
                class="${properties.kcButtonClass!} ${properties.kcButtonPrimaryClass!} ${properties.kcButtonBlockClass!} ${properties.kcButtonLargeClass!}"
                name="login" id="kc-login" type="submit" value="${msg(" doLogIn")}" />
            </div>
            -->

            <div>
                <div>
                    <button tabindex="0" name="login" id="kc-login" type="submit" class="submitbutton">
                        ${msg("doLogIn")}
                    </button>
                </div>
            </div>

            </form>
            </#if>
            </div>
            <#if realm.password && social.providers??>
                <div id="kc-social-providers"
                    class="${properties.kcFormSocialAccountContentClass!} ${properties.kcFormSocialAccountClass!}">
                    <ul
                        class="${properties.kcFormSocialAccountListClass!} <#if social.providers?size gt 4>${properties.kcFormSocialAccountDoubleListClass!}</#if>">
                        <#list social.providers as p>
                            <li class="${properties.kcFormSocialAccountListLinkClass!}"><a href="${p.loginUrl}"
                                    id="zocial-${p.alias}" class="zocial ${p.providerId}">
                                    <span>${p.displayName}</span></a></li>
                        </#list>
                    </ul>
                </div>
            </#if>
            </div>
            <#elseif section="info">
                <#if realm.password && realm.registrationAllowed && !registrationDisabled??>
                    <hr class="mdc-list-divider divider-mdc">
                    <div class="legal-lttr">
                        Created by Listen To The Rainbow@2021 LttR holding all rights reserved
                    </div>
                </#if>
            </#if>



    </@layout.registrationLayout>