// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tData struct {}
var Data tData


func (_ tData) SaveServer(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.SaveServer", args).Url
}

func (_ tData) DeleteServer(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.DeleteServer", args).Url
}

func (_ tData) LoadServers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.LoadServers", args).Url
}

func (_ tData) LoadSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.LoadSettings", args).Url
}

func (_ tData) SaveSettings(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.SaveSettings", args).Url
}

func (_ tData) FindPackets(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.FindPackets", args).Url
}

func (_ tData) CountPackets(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Data.CountPackets", args).Url
}


type tShows struct {}
var Shows tShows


func (_ tShows) Load(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.Load", args).Url
}

func (_ tShows) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.Save", args).Url
}

func (_ tShows) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.Delete", args).Url
}

func (_ tShows) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.Search", args).Url
}

func (_ tShows) LoadEpisodes(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.LoadEpisodes", args).Url
}

func (_ tShows) RecentEpisodes(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Shows.RecentEpisodes", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


