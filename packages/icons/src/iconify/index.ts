import { addCollection, createIconifyIcon } from '@vben-core/icons';

import emojioneMonotone from './emojione-monotone.json';
import skillIcons from './skill-icons.json';

addCollection(emojioneMonotone);
addCollection(skillIcons);

export * from '@vben-core/icons';

export const MdiKeyboardEsc = createIconifyIcon('mdi:keyboard-esc');

export const MdiWechat = createIconifyIcon('mdi:wechat');

export const MdiGithub = createIconifyIcon('mdi:github');

export const MdiGoogle = createIconifyIcon('mdi:google');

export const MdiQqchat = createIconifyIcon('mdi:qqchat');
