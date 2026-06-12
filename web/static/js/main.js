const THEME_KEY = "recoverease-theme";

const getPreferredTheme = () => {
  const savedTheme = localStorage.getItem(THEME_KEY);

  if (savedTheme === "light" || savedTheme === "dark") {
    return savedTheme;
  }

  return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
};

const applyTheme = (theme) => {
  const nextTheme = theme === "dark" ? "dark" : "light";
  const themeButton = document.querySelector(".theme-toggle");

  document.body.dataset.theme = nextTheme;

  if (themeButton) {
    const isDark = nextTheme === "dark";
    themeButton.setAttribute("aria-pressed", String(isDark));
    themeButton.setAttribute("aria-label", isDark ? "Switch to light theme" : "Switch to dark theme");
    themeButton.setAttribute("title", isDark ? "Switch to light theme" : "Switch to dark theme");
  }

  localStorage.setItem(THEME_KEY, nextTheme);
};

document.querySelectorAll('a[href^="#"]').forEach((link) => {
  link.addEventListener("click", (event) => {
    const target = document.querySelector(link.getAttribute("href"));

    if (!target) {
      return;
    }

    event.preventDefault();
    target.scrollIntoView({ behavior: "smooth", block: "start" });
  });
});

document.querySelectorAll(".locale-option").forEach((option) => {
  option.addEventListener("click", (event) => {
    if (!window.location.hash) {
      return;
    }

    event.preventDefault();
    window.location.href = `${option.getAttribute("href")}${window.location.hash}`;
  });
});

document.addEventListener("click", (event) => {
  document.querySelectorAll(".locale-menu[open]").forEach((menu) => {
    if (!menu.contains(event.target)) {
      menu.removeAttribute("open");
    }
  });
});

document.querySelector(".theme-toggle")?.addEventListener("click", () => {
  const currentTheme = document.body.dataset.theme === "dark" ? "dark" : "light";
  applyTheme(currentTheme === "dark" ? "light" : "dark");
});

applyTheme(getPreferredTheme());
