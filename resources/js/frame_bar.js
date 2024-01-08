document.writeln("<link rel=\'stylesheet\' href=\'css/bulma.css\'>");
document.writeln("<link rel=\'stylesheet\' href=\'css/main.css\'>");
document.writeln("<script type=\'text/javascript\' src=\'js/utils.js\'></script>");
document.writeln("<script src=\'https://kit.fontawesome.com/fde73c65f6.js\' crossorigin=\'anonymous\'></script>");
document.writeln("<nav class=\'navbar has-shadow is-fixed-top\' role=\'navigation\' aria-label=\'frame bar\'>");
document.writeln("    <div class=\'navbar-brand\' style=\'-webkit-app-region: drag;\'>");
document.writeln("        <a class=\'navbar-item\'>");
document.writeln("            <span class=\'icon-text has-text-info is-size-5\'>");
document.writeln("                <span><strong>吃了吗</strong></span>");
document.writeln("                <span>");
document.writeln("                    <i class=\'fas fa-utensils \'></i>");
document.writeln("                </span>");
document.writeln("            </span>");
document.writeln("        </a>");
document.writeln("    </div>");
document.writeln("    <div class=\'navbar-menu is-active\'>");
document.writeln("        <div class=\'navbar-end\'>");
document.writeln("            <div class=\'navbar-item\'>");
document.writeln("                <div class=\'buttons\'>");
document.writeln("                    <a class=\'button is-success is-size-7 is-outlined\' onclick=\'onMinWindow()\'");
document.writeln("                        style=\'font-family: Marlett24;\'>");
document.writeln("                        <strong style=\'font-size: large;\'>0</strong>");
document.writeln("                    </a>");
document.writeln("                </div>");
document.writeln("            </div>");
document.writeln("            <div class=\'navbar-item\'>");
document.writeln("                <div class=\'buttons\'>");
document.writeln("                    <a class=\'button is-danger is-size-7 is-outlined\' onclick=\'onCloseWindow()\'");
document.writeln("                        style=\'font-family: Marlett24;\'>");
document.writeln("                        <strong style=\'font-size: large;\'>r</strong>");
document.writeln("                    </a>");
document.writeln("                </div>");
document.writeln("            </div>");
document.writeln("        </div>");
document.writeln("    </div>");
document.writeln("</nav></br>");
document.writeln("");
document.writeln("<nav class=\'navbar is-fixed-bottom is-light\' role=\'navigation\' aria-label=\'main navigation\'>");
document.writeln("    <div class=\'navbar-brand\'>");
document.writeln("        <a role=\'button\' class=\'navbar-burger\' aria-label=\'menu\' aria-expanded=\'false\' data-target=\'navbar-main\'>");
document.writeln("            <span aria-hidden=\'true\'></span>");
document.writeln("            <span aria-hidden=\'true\'></span>");
document.writeln("            <span aria-hidden=\'true\'></span>");
document.writeln("        </a>");
document.writeln("    </div>");
document.writeln("");
document.writeln("    <div id=\'navbar-main\' class=\'navbar-menu\'>");
document.writeln("        <div class=\'navbar-start\'>");
document.writeln("            <a class=\'navbar-item\' href=\'index.html\'>");
document.writeln("                首页");
document.writeln("            </a>");
document.writeln("            <a class=\'navbar-item\' href=\'about.html\'>");
document.writeln("                关于");
document.writeln("            </a>");
document.writeln("            <div class=\'navbar-item has-dropdown has-dropdown-up is-hoverable\'>");
document.writeln("                <a class=\'navbar-link\'>");
document.writeln("                    功能");
document.writeln("                </a>");
document.writeln("                <div class=\'navbar-dropdown\'>");
document.writeln("                    <a class=\'navbar-item\' href=\'unlock_music.html\'>");
document.writeln("                        音乐解锁");
document.writeln("                    </a>");
document.writeln("                    <hr class=\'navbar-divider\'>");
document.writeln("                    <a class=\'navbar-item\' onclick=\'onCloseWindow()\'>");
document.writeln("                        关闭软件");
document.writeln("                    </a>");
document.writeln("                </div>");
document.writeln("            </div>");
document.writeln("        </div>");
document.writeln("");
document.writeln("        <div class=\'navbar-end\'>");
document.writeln("            <div class=\'navbar-item\'>");
document.writeln("                <div class=\'buttons\'>");
document.writeln("                    <a class=\'button is-info\' onclick=\'onContactUs()\'>");
document.writeln("                        <strong>联系我们</strong>");
document.writeln("                    </a>");
document.writeln("                </div>");
document.writeln("            </div>");
document.writeln("        </div>");
document.writeln("    </div>");
document.writeln("    <script>");
document.writeln("        document.addEventListener(\'DOMContentLoaded\', () => {");
document.writeln("            // Get all \'navbar-burger\' elements");
document.writeln("            const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll(\'.navbar-burger\'), 0);");
document.writeln("            // Add a click event on each of them");
document.writeln("            $navbarBurgers.forEach(el => {");
document.writeln("                el.addEventListener(\'click\', () => {");
document.writeln("");
document.writeln("                    // Get the target from the \'data-target\' attribute");
document.writeln("                    const target = el.dataset.target;");
document.writeln("                    const $target = document.getElementById(target);");
document.writeln("");
document.writeln("                    // Toggle the \'is-active\' class on both the \'navbar-burger\' and the \'navbar-menu\'");
document.writeln("                    el.classList.toggle(\'is-active\');");
document.writeln("                    $target.classList.toggle(\'is-active\');");
document.writeln("");
document.writeln("                });");
document.writeln("            });");
document.writeln("        });");
document.writeln("    </script>");
document.writeln("</nav>");