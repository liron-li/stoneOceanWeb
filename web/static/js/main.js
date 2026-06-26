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

document.querySelectorAll('input[name="license"]').forEach((option) => {
  option.addEventListener("change", () => {
    const price = Number(option.dataset.price || 0);
    const formattedPrice = `$${price.toFixed(2)}`;

    document.querySelectorAll("[data-license-price]").forEach((target) => {
      target.textContent = formattedPrice;
    });
  });
});

const isChinesePage = document.documentElement.lang.toLowerCase().startsWith("zh");
const checkoutButton = document.querySelector(".checkout-submit");
const recoveryButton = document.querySelector(".license-recovery-submit");

const text = {
  missingEmail: isChinesePage ? "请先填写邮箱地址。" : "Enter your email address first.",
  checkoutPending: isChinesePage ? "正在创建订单并确认测试支付..." : "Creating order and confirming test payment...",
  checkoutSuccess: isChinesePage ? "支付已确认，激活码已生成。" : "Payment confirmed. Your license key is ready.",
  recoveryPending: isChinesePage ? "正在查询这个邮箱的激活码..." : "Looking up license keys for this email...",
  recoveryEmpty: isChinesePage ? "没有找到这个邮箱的购买记录。" : "No purchases were found for this email.",
  recoverySuccess: isChinesePage ? "已找到以下激活码。" : "License keys found.",
  requestFailed: isChinesePage ? "请求失败，请稍后重试。" : "Request failed. Please try again.",
  orderNo: isChinesePage ? "订单号" : "Order",
  licenseKey: isChinesePage ? "激活码" : "License key",
  expiresAt: isChinesePage ? "有效期至" : "Expires",
  lifetime: isChinesePage ? "永久有效" : "Lifetime",
};

const postJSON = async (url, body = {}) => {
  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  const data = await response.json().catch(() => ({}));
  if (!response.ok) {
    throw new Error(data.error || text.requestFailed);
  }
  return data;
};

const setStatusText = (target, type, message) => {
  if (!target) {
    return;
  }
  target.className = `form-status ${type}`;
  target.textContent = message;
};

const renderLicenses = (target, message, licenses) => {
  if (!target) {
    return;
  }

  target.className = "form-status success";
  target.textContent = "";

  const intro = document.createElement("p");
  intro.textContent = message;
  target.append(intro);

  const list = document.createElement("div");
  list.className = "license-result-list";

  licenses.forEach((license) => {
    const item = document.createElement("article");
    item.className = "license-result";

    const key = document.createElement("strong");
    key.textContent = license.licenseKey;

    const order = document.createElement("span");
    order.textContent = `${text.orderNo}: ${license.orderNo}`;

    const expiry = document.createElement("span");
    expiry.textContent = license.expiresAt
      ? `${text.expiresAt}: ${new Date(license.expiresAt).toLocaleDateString()}`
      : text.lifetime;

    item.append(key, order, expiry);
    list.append(item);
  });

  target.append(list);
};

checkoutButton?.addEventListener("click", async () => {
  const status = document.querySelector(".checkout-status");
  const email = document.querySelector('input[name="email"]')?.value.trim();
  const license = document.querySelector('input[name="license"]:checked')?.value;
  const paymentMethod = document.querySelector('input[name="payment"]:checked')?.value;

  if (!email) {
    setStatusText(status, "error", text.missingEmail);
    return;
  }

  checkoutButton.disabled = true;
  setStatusText(status, "pending", text.checkoutPending);

  try {
    const checkout = await postJSON("/api/checkout", { email, license, paymentMethod });
    const payment = await postJSON(checkout.paymentUrl);
    renderLicenses(status, text.checkoutSuccess, [payment.license]);
  } catch (error) {
    setStatusText(status, "error", error.message || text.requestFailed);
  } finally {
    checkoutButton.disabled = false;
  }
});

recoveryButton?.addEventListener("click", async () => {
  const status = document.querySelector(".recovery-status");
  const email = document.querySelector('input[name="recovery-email"]')?.value.trim();

  if (!email) {
    setStatusText(status, "error", text.missingEmail);
    return;
  }

  recoveryButton.disabled = true;
  setStatusText(status, "pending", text.recoveryPending);

  try {
    const result = await postJSON("/api/license-recovery", { email });
    if (!result.licenses || result.licenses.length === 0) {
      setStatusText(status, "empty", text.recoveryEmpty);
      return;
    }
    renderLicenses(status, text.recoverySuccess, result.licenses);
  } catch (error) {
    setStatusText(status, "error", error.message || text.requestFailed);
  } finally {
    recoveryButton.disabled = false;
  }
});

applyTheme(getPreferredTheme());
