<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let listCateitem = [];
    let search = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;
    let totalrecordall = 0;
    let totalpaging = 0;
    let perpage = 0;
    let page = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "CURR-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome(e) {
        listHome = [];
        const res = await fetch("/api/item", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                item_search: e,
                item_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listcateitem = json.listcateitem;
            record_message = json.message;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["item_id"],
                            home_idmerek: record[i]["item_idmerek"],
                            home_nmmerek: record[i]["item_nmmerek"],
                            home_idcateitem: record[i]["item_idcateitem"],
                            home_nmcateitem: record[i]["item_nmcateitem"],
                            home_iduom: record[i]["item_iduom"],
                            home_name: record[i]["item_name"],
                            home_descp: record[i]["item_descp"],
                            home_urlimg: record[i]["item_urlimg"],
                            home_inventory: record[i]["item_inventory"],
                            home_sales: record[i]["item_sales"],
                            home_purchase: record[i]["item_purchase"],
                            home_status: record[i]["item_status"],
                            home_status_css: record[i]["item_status_css"],
                            home_sales_css: record[i]["item_sales_css"],
                            home_purchase_css: record[i]["item_purchase_css"],
                            home_inventory_css: record[i]["item_inventory_css"],
                            home_create: record[i]["item_create"],
                            home_update: record[i]["item_update"],
                        },
                    ];
                }
                listPage = [];
                for(var i=1;i<=totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listcateitem.length; i++) {
                listCateitem = [
                    ...listCateitem,
                    {
                        cateitem_id: record_listcateitem[i]["cateitem_id"],
                        cateitem_name: record_listcateitem[i]["cateitem_name"],
                    },
                ];
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        listCateitem = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    const handleSearch = (e) => {
        search = e.detail.searchHome;
        initHome(search,"")
    };
    const handlePaging = (e) => {
        page = e.detail.page
        initHome("")
    };
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleSearch={handleSearch}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listPage}
    {listHome}
    {listCateitem}
    {totalrecord}/>
{/if}