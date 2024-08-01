<template>
  <div ref="rootElement" class="user-bar" v-click-outside="onClickOutside">
    <div v-if="auth.user?.nickname" class="user-bar-name">
      {{ auth.user.nickname }}
    </div>
    <ElAvatar :size="50" :src="avatarUrl" />
    <ElPopover
      ref="popoverRef"
      trigger="click"
      virtual-triggering
      persistent
      :virtual-ref="rootElement"
    >
      <div class="user-bar-popover-menu-item" @click="onClickProfile">
        <ElIcon>
          <IEpStar />
        </ElIcon>
        <span>我的</span>
      </div>
      <div class="user-bar-popover-menu-separator"></div>
      <div class="user-bar-popover-menu-item" @click="onClickSignOut">
        <ElIcon>
          <IEpSwitchButton />
        </ElIcon>
        <span>登出</span>
      </div>
    </ElPopover>
  </div>
</template>

<script lang="ts" setup>
import { ref, unref, computed } from "vue";
import { ClickOutside as vClickOutside } from "element-plus";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/AuthStore";

const rootElement = ref();
const popoverRef = ref();

const router = useRouter();
const auth = useAuthStore();

const avatarUrl = computed(() => {
  return auth.user?.avatar_path ?? "/avatar-circle.png";
});

const onClickOutside = () => {
  unref(popoverRef).popperRef?.delayHide?.();
};

const onClickProfile = () => {
  router.push("/setting/profile");
};

const onClickSignOut = () => {
  auth.token = "";
  router.replace("/login");
};
</script>

<style lang="scss" scoped>
.user-bar {
  --avatar-size: 50px;

  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.user-bar-name {
  padding: 0 1em;
}

.user-bar-popover-menu-item {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 1.4vw;

  &:hover {
    color: #47f;
  }

  // & > span {
  //   flex-grow: 1;
  // }
}

.user-bar-popover-menu-separator {
  margin: .5ex 0;
  border-bottom: 1px solid #ddd;
}
</style>
