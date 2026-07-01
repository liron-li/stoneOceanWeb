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

const buildLocaleHref = (href) => {
  const target = new URL(href, window.location.origin);
  if (document.querySelector("[data-payment-result]")) {
    new URLSearchParams(window.location.search).forEach((value, key) => {
      if (!target.searchParams.has(key)) {
        target.searchParams.append(key, value);
      }
    });
  }
  target.hash = window.location.hash;
  return `${target.pathname}${target.search}${target.hash}`;
};

document.querySelectorAll(".locale-option").forEach((option) => {
  option.addEventListener("click", (event) => {
    const shouldPreservePaymentQuery = document.querySelector("[data-payment-result]") && window.location.search;
    if (!window.location.hash && !shouldPreservePaymentQuery) {
      return;
    }

    event.preventDefault();
    window.location.href = buildLocaleHref(option.getAttribute("href"));
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
const pageLocale = document.documentElement.lang || navigator.language || "en";
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

const getJSON = async (url) => {
  const response = await fetch(url);
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

const formatLocalizedDate = (value, options) => {
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return value || "";
  }
  return new Intl.DateTimeFormat(pageLocale, options).format(date);
};

const formatDate = (value) => formatLocalizedDate(value, { dateStyle: "medium" });
const formatDateTime = (value) => formatLocalizedDate(value, { dateStyle: "medium", timeStyle: "short" });

const focusInvalidEmail = (input) => {
  if (!input) {
    return;
  }

  input.setAttribute("aria-invalid", "true");
  input.scrollIntoView({ behavior: "smooth", block: "center" });
  input.focus({ preventScroll: true });
};

document.querySelectorAll('input[type="email"]').forEach((input) => {
  input.addEventListener("input", () => {
    input.removeAttribute("aria-invalid");
  });
});

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
      ? `${text.expiresAt}: ${formatDate(license.expiresAt)}`
      : text.lifetime;

    item.append(key, order, expiry);
    list.append(item);
  });

  target.append(list);
};

checkoutButton?.addEventListener("click", async () => {
  const status = document.querySelector(".checkout-status");
  const emailInput = document.querySelector('input[name="email"]');
  const email = emailInput?.value.trim();
  const license = document.querySelector('input[name="license"]:checked')?.value;
  const paymentMethod = document.querySelector('input[name="payment"]:checked')?.value;

  if (!email) {
    setStatusText(status, "error", text.missingEmail);
    focusInvalidEmail(emailInput);
    return;
  }

  checkoutButton.disabled = true;
  setStatusText(status, "pending", text.checkoutPending);

  try {
    const checkout = await postJSON("/api/checkout", { email, license, paymentMethod });
    await postJSON(checkout.paymentUrl);

    const successPath = checkoutButton.dataset.successPath || "/checkout/success";
    window.location.href = `${successPath}?paymentNo=${encodeURIComponent(checkout.paymentNo)}`;
  } catch (error) {
    setStatusText(status, "error", error.message || text.requestFailed);
  } finally {
    checkoutButton.disabled = false;
  }
});

recoveryButton?.addEventListener("click", async () => {
  const status = document.querySelector(".recovery-status");
  const emailInput = document.querySelector('input[name="recovery-email"]');
  const email = emailInput?.value.trim();

  if (!email) {
    setStatusText(status, "error", text.missingEmail);
    focusInvalidEmail(emailInput);
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

const copyText = async (value) => {
  if (navigator.clipboard?.writeText) {
    await navigator.clipboard.writeText(value);
    return;
  }

  const input = document.createElement("textarea");
  input.value = value;
  input.setAttribute("readonly", "");
  input.style.position = "fixed";
  input.style.opacity = "0";
  document.body.append(input);
  input.select();
  document.execCommand("copy");
  input.remove();
};

const renderPaymentResult = async () => {
  const root = document.querySelector("[data-payment-result]");
  if (!root) {
    return;
  }

  const status = root.querySelector("[data-payment-status]");
  const panel = root.querySelector("[data-license-panel]");
  const params = new URLSearchParams(window.location.search);
  const paymentNo = params.get("paymentNo");

  if (!paymentNo) {
    status.textContent = root.dataset.missing;
    status.className = "payment-result-status error";
    return;
  }

  status.textContent = root.dataset.loading;
  status.className = "payment-result-status pending";

  try {
    const result = await getJSON(`/api/payments/${encodeURIComponent(paymentNo)}/result`);
    if (!result.license) {
      status.textContent = root.dataset.pending;
      status.className = "payment-result-status pending";
      return;
    }

    const license = result.license;
    status.textContent = "";
    status.className = "payment-result-status success";
    panel.hidden = false;
    panel.textContent = "";

    const keyRow = document.createElement("div");
    keyRow.className = "license-key-row";

    const key = document.createElement("strong");
    key.textContent = license.licenseKey;

    const copyButton = document.createElement("button");
    copyButton.className = "button button-secondary license-copy-button";
    copyButton.type = "button";
    copyButton.textContent = root.dataset.copy;
    copyButton.addEventListener("click", async () => {
      await copyText(license.licenseKey);
      copyButton.textContent = root.dataset.copied;
      window.setTimeout(() => {
        copyButton.textContent = root.dataset.copy;
      }, 1800);
    });

    keyRow.append(key, copyButton);

    const detailList = document.createElement("dl");
    detailList.className = "license-detail-list";

    const details = [
      [root.dataset.order, license.orderNo || result.orderNo],
      [root.dataset.plan, license.plan],
      [root.dataset.issuedAt, formatDateTime(license.issuedAt)],
      [
        root.dataset.expiresAt,
        license.expiresAt ? formatDate(license.expiresAt) : root.dataset.lifetime,
      ],
    ];

    details.forEach(([label, value]) => {
      const row = document.createElement("div");
      const term = document.createElement("dt");
      const desc = document.createElement("dd");
      term.textContent = label;
      desc.textContent = value;
      row.append(term, desc);
      detailList.append(row);
    });

    panel.append(keyRow, detailList);
  } catch (error) {
    status.textContent = error.message || root.dataset.failed;
    status.className = "payment-result-status error";
  }
};

renderPaymentResult();

applyTheme(getPreferredTheme());
