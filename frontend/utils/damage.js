import CPJalaliDate from './JalaliDate';

export const damageTypes = {
  death_damage: 'فوت',
  lack_damage: 'نقص عضو',
  medical_damage: 'پزشکی',
  explosion_damage: 'انفجار',
  instrument_damage: 'تجهیزات',
  firing_damage: 'آتش سوزی',
};
export const NotCoveredReasonTypes = {
  1: 'علي رغم اعلام به مشترک در سامانه بيمه، مراجعه حضوري جهت ارائه مدارک تکميلي انجام نشده است.',
  2: 'تجهيز حادثه ديده، جزء تجهيزات تحت پوشش بيمه نيست.',
  3: 'فراواني حوادث واقع شده براي مشترکين يک منطقه، نشان دهنده مشکلي از جانب شبکه است.',
  21: 'دوفاز شدن شبکه در پي علل مختلف، مسبب اين حادثه بوده است.',
  41: 'پرونده تکراری است.',
};

export const missingDocumentTypes = {
  '2000': 'اصلاح اطلاعات بانکی و شبا',
  '3': 'آخرين قبض پرداختي',
  '2': 'تصوير کارت ملي',
  '1': 'سند مالکيت يا اجاره نامه',
  '41': 'تصوير فاکتور فروشگاه يا تعميرگاه مجاز',
  '49': 'تصوير گزارش معتمدين محل يا نيروي انتظامي',
  '51': 'تصوير گزارش کلانتري',
  '52': 'تصوير پرونده دادگاه در صورت شکايت',
  '32': 'صورت حساب پزشکي پرونده مصدوم',
  '50': 'تصوير گزارش آتش نشاني',
  '31': 'گواهي پزشک معالج',
  '39': 'تصوير محل حادثه',
  '30': 'کپي کارت ملي، شناسنامه مصدوم',
  '38': 'تصوير فاکتور تعميرات',
  '29': 'اولين مرجع درماني پزشکي معالج',
  '37': 'تصوير ليست موارد آسيب ديده',
  '28': 'راديولوژي بعد از حادثه',
  '34': 'تصوير مدارک پزشکي و پرونده هاي بيمارستاني',
  '36': 'تصوير گزارش آتش نشاني يا مقامات ذيصلاح براساس نوع انفجار',
  '33': 'تصوير اصل صورت حساب بيمارستان',
  '27': 'گزارش انتظامي يا بازرس آتش نشاني',
  '21': 'تصوير راي قاضي',
  '22': 'گواهي فوت',
  '35': 'تصوير گزارش مقامات ذيصلاح',
  '23': 'شناسنامه ابطال شده، کارت ملي',
  '24': 'گزارش پزشک قانوني',
  '25': 'تصوير معاينه جسد',
  '26': 'گواهي انحصار وراثت',
};

export const tavanirDamageTypes = {
  '1': 'فوت',
  '2': 'نقص عضو',
  '3': 'هزينه پزشکي',
  '4': 'انفجار',
  '5': 'لوازم تجهيزات',
  '6': 'آتش سوزي',
};

export const tavanirDamageMaxAmount = {
  '1': 9778200000,
  '2': 6518800000,
  '3': 6926228000,
  '4': 8963350000,
  '5': 260752000,
  '6': 8963350000,
};

export const statusChoices = {
  NEW: 'جدید',
  ACCEPTED: 'تایید شده',
  AMOUNT_REJECTED: 'مبلغ رد شده',
  CANCELED_BY_USER: 'انصراف ذی نفع',
  CLOSED: 'مختومه',
  COMPLETED: 'ارسال اولیه',
  NOT_COMPLETED: 'تکمیل نشده',
  IN_PROGRESS: 'جاری',
  INACTIVE: 'غیر فعال',
  INCOMPLETE: 'درخواست ناقص',
  INCOMPLETE_CHANGE: 'تکمیل درخواست ناقص',
  PAYED: 'پرداخت شده',
  READY_TO_PAY: 'آماده پرداخت',
  SUSPENDED: 'معوق',
  REQUEST_COMPLIANT: 'اعتراض مشترک',

};

export const tavanirStatusChoices = {
  NEW: 'جدید',
  CLOSED: 'مختومه',
  IN_PROGRESS: 'جاری',
  INCOMPLETE: 'درخواست ناقص',
  INCOMPLETE_CHANGE: 'تکمیل درخواست ناقص',
  PAYED: 'پرداخت شده',
  READY_TO_PAY: 'آماده پرداخت',
  SUSPENDED: 'معوق',
  REJECTED: 'رد شده',
  REQUEST_COMPLIANT: 'اعتراض مشترک',
};

export const expertStatusChoices = {
  DEFAULT: 'در حال کارشناسی',
  REJECTED: 'رد شده',
  INCOMPLETE: 'ناقص',
  INCOMPLETE_CHANGE: 'تکمیل درخواست ناقص',
  PAYED: 'پرداخت شده',
  READY_TO_PAY: 'آماده پرداخت',
  Need_Visit_InPerson: 'نیاز به مراجعه حضوری',
  REQUEST_COMPLIANT: 'اعتراض مشترک',
};

export const tavanirMessageType = {
  assigned: 'assigned',
  documentUpdate: 'documentUpdate',
  tavanirCoverCompliant: 'tavanirCoverCompliant',
  amountCompliant: 'amountCompliant',
  tavanirAmountCompliant: 'tavanirAmountCompliant',
  updateBankInfo: 'updateBankInfo',
  checkoutCompliant: 'checkoutCompliant',
};

export const getStatusChoicesTitle = f => {
  return statusChoices[f] || expertStatusChoices[f] || f;
};
