<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./components/Navbar.svelte";
	import Dashboard from "./pages/dashboard/Dashboard.svelte";
	
	import Admin from "./pages/admin/Admin.svelte";
	import Adminrule from "./pages/adminrule/Adminrule.svelte";
	import Currency from "./pages/currency/Currency.svelte";
	import Employee from "./pages/employee/Employee.svelte";
	import Departement from "./pages/departement/Departement.svelte";
	import Catevendor from "./pages/catevendor/Catevendor.svelte";
	import Vendor from "./pages/vendor/Vendor.svelte";
	import Uom from "./pages/uom/Uom.svelte";
	import Branch from "./pages/branch/Branch.svelte";
	import Warehouse from "./pages/warehouse/Warehouse.svelte";
	import Merek from "./pages/merek/Merek.svelte";
	import Categoryitem from "./pages/categoryitem/Categoryitem.svelte";
	import Purchaserequest from "./pages/purchaserequest/Purchaserequest.svelte";
	import Rfq from "./pages/rfq/Rfq.svelte";
	import Item from "./pages/item/Item.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/NotFound.svelte";
	export let table_header_font;
	export let table_body_font;
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
				component: Dashboard,
			}),
			"/rfq": wrap({
				component: Rfq,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/merek": wrap({
				component: Merek,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/purchaserequest": wrap({
				component: Purchaserequest,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/employee": wrap({
				component: Employee,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/departement": wrap({
				component: Departement,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/catevendor": wrap({
				component: Catevendor,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/vendor": wrap({
				component: Vendor,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/item": wrap({
				component: Item,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/categoryitem": wrap({
				component: Categoryitem,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/listuom": wrap({
				component: Uom,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/listbranch": wrap({
				component: Branch,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/listwarehouse": wrap({
				component: Warehouse,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/currency": wrap({
				component: Currency,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			
			"/admin": wrap({
				component: Admin,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/adminrule": wrap({
				component: Adminrule,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navbar />
{/if}
<Router {routes} />
