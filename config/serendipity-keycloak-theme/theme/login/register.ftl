<#import "template.ftl" as layout>
    <@layout.registrationLayout; section>
        <#if section="header">
            ${msg("registerTitle")}
            <#elseif section="form">
                <form id="kc-register-form" class="${properties.kcFormClass!}" action="${url.registrationAction}"
                    method="post">

                    <div>
                        <div>
                            <label for="firstName" class="inputlabel">${msg("firstName")}</label>
                        </div>
                        <div>
                            <input type="text" id="firstName" class="inputs ${properties.kcInputClass!}"
                                name="firstName" value="${(register.formData.firstName!'')}" />
                        </div>
                    </div>



                    <div>
                        <div>
                            <label for="lastName" class="inputlabel">${msg("lastName")}</label>
                        </div>
                        <div>
                            <input type="text" id="lastName" class="inputs ${properties.kcInputClass!}" name="lastName"
                                value="${(register.formData.lastName!'')}" />
                        </div>
                    </div>


                    <div>
                        <div>
                            <label for="email" class="inputlabel">${msg("email")}</label>
                        </div>
                        <div>
                            <input type="text" id="email" class="inputs ${properties.kcInputClass!}" name="email"
                                value="${(register.formData.email!'')}" autocomplete="email" />
                        </div>
                    </div>


                    <#if !realm.registrationEmailAsUsername>
                        <div>
                            <div>
                                <label for="username" class="inputlabel">${msg("username")}</label>
                            </div>
                            <div>
                                <input type="text" id="username" class="inputs ${properties.kcInputClass!}"
                                    name="username" value="${(register.formData.username!'')}"
                                    autocomplete="username" />
                            </div>
                        </div>
                    </#if>

                    <#if passwordRequired??>
                        <div>
                            <div>
                                <label for="password" class="inputlabel">${msg("password")}</label>
                            </div>
                            <div>
                                <input type="password" id="password" class="inputs ${properties.kcInputClass!}"
                                    name="password" autocomplete="new-password" />
                            </div>
                        </div>

                        <div>
                            <div>
                                <label for="password-confirm" class="inputlabel">${msg("passwordConfirm")}</label>
                            </div>
                            <div>
                                <input type="password" id="password-confirm" class="inputs ${properties.kcInputClass!}"
                                    name="password-confirm" />
                            </div>
                        </div>
                    </#if>

                    <div class="${properties.kcFormGroupClass!}">
                        <div id="kc-form-options" class="${properties.kcFormOptionsClass!}">
                            <div class="${properties.kcFormOptionsWrapperClass!}">
                                <span><a href="${url.loginUrl}">${kcSanitize(msg("backToLogin"))?no_esc}</a></span>
                            </div>
                        </div>

                        <div id="kc-form-buttons">
                            <input class="submitbutton" type="submit" value="${msg("doRegister")}" />
                        </div>
                    </div>

                    <div class="mdc-card__actions">

                        <#-- <button class="mdc-button mdc-card__action mdc-card__action--button"
                            onclick="window.location.href = ${url.loginUrl}">
                            <i class="material-icons mdc-button__icon">arrow_back</i>
                            ${kcSanitize(msg("backToLogin"))?no_esc}
                            </button>
                            -->

                            <a href="${url.loginUrl}" class="mdc-button mdc-card__action mdc-card__action--button">
                                <i class="material-icons mdc-button__icon">arrow_back</i>
                                ${kcSanitize(msg("backToLogin"))?no_esc}
                            </a>
                            <!-- 
                            <div class="mdc-card__action-icons">
                                <div class="mdc-card__action-buttons">
                                    <button tabindex="0" name="login" id="kc-login" type="submit"
                                        class="mdc-button mdc-button--raised mdc-card__action">
                                        ${msg("doRegister")}
                                    </button>
                                </div>
                            </div> -->

                    </div>

                </form>
        </#if>
        </@layout.registrationLayout>