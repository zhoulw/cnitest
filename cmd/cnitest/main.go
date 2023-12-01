package main

import (
	"cnitest/cni"
	"cnitest/helper"
	"cnitest/utils"
	"errors"
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/utils/buildversion"
)

const (
	pluginName   = "cnitest"
	buildVersion = "v0.3.0"
)

func cmdAdd(args *skel.CmdArgs) error {
	utils.WriteLog("begin to cni add·········")

	helper.TmpLogArgs(args)

	// 从 args 里把 config 给捞出来
	pluginConfig := helper.GetConfigs(args)
	if pluginConfig == nil {
		errMsg := fmt.Sprintf("add: 从 args 中获取 plugin config 失败, config: %s", string(args.StdinData))
		utils.WriteLog(errMsg)
		return errors.New(errMsg)
	}

	mode, cniVersion := helper.GetBaseInfo(pluginConfig)
	if pluginConfig.CNIVersion == "" {
		pluginConfig.CNIVersion = cniVersion
	}

	// 将 args 和 configs 以及要使用的插件模式都传给 cni manager
	cniManager := cni.GetCNIManager().
		SetBootstrapConfigs(pluginConfig).
		SetBootstrapArgs(args).
		SetBootstrapCNIMode(mode)
	if cniManager == nil {
		utils.WriteLog("cni 插件未初始化完成")
		return errors.New("cni plugins register failed")
	}

	// 启动对应 mode 的插件开始设置乱七八糟的网卡等
	err := cniManager.BootstrapCNI()
	if err != nil {
		utils.WriteLog("设置 cni 失败: ", err.Error())
		return err
	}

	// 将结果打印到标准输出
	err = cniManager.PrintResult()
	if err != nil {
		utils.WriteLog("打印 cni 执行结果失败: ", err.Error())
		return err
	}

	return nil
}

func cmdDel(args *skel.CmdArgs) error {

	utils.WriteLog("进入到 cmdDel")
	helper.TmpLogArgs(args)

	pluginConfig := helper.GetConfigs(args)
	if pluginConfig == nil {
		errMsg := fmt.Sprintf("del: 从 args 中获取 plugin config 失败, config: %s", string(args.StdinData))
		utils.WriteLog(errMsg)
		return errors.New(errMsg)
	}
	mode, _ := helper.GetBaseInfo(pluginConfig)

	cniManager := cni.
		GetCNIManager().
		SetUnmountConfigs(pluginConfig).
		SetUnmountArgs(args).
		SetUnmountCNIMode(mode)

	// 这里的 del 如果返回 error 的话, kubelet 就会尝试一直不停地执行 StopPodSandbox
	// 直到删除后的 error 返回 nil 未知
	// return errors.New("test cmdDel")
	return cniManager.UnmountCNI()
}

func cmdCheck(args *skel.CmdArgs) error {

	utils.WriteLog("brgin to cni check········")
	utils.WriteLog(
		"containerID: ", args.ContainerID,
		"netns:", args.Netns,
		"IfName: ", args.IfName,
		"Path: ", args.Path,
		"StdinData: ", string(args.StdinData),
	)
	return nil
}

func main() {
	buildversion.BuildVersion = buildVersion

	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, buildversion.BuildString(pluginName))
}
