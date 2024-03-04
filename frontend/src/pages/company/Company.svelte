<script>
    import Home from "./Home.svelte";
   
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let listCurrency = [];
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
                page: "COMPANY-VIEW",
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
        const res = await fetch("/api/company", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_search: e,
                company_page : parseInt(page)
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            let record_listcurr = json.listcurr;
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
                            home_id: record[i]["company_id"],
                            home_idcurr: record[i]["company_idcurr"],
                            home_startjoin: record[i]["company_startjoin"],
                            home_endjoin: record[i]["company_endjoin"],
                            home_name: record[i]["company_name"],
                            home_owner: record[i]["company_owner"],
                            home_email: record[i]["company_email"],
                            home_hp1: record[i]["company_phone1"],
                            home_hp2: record[i]["company_phone2"],
                            home_minfee: record[i]["company_minfee"],
                            home_url1: record[i]["company_url1"],
                            home_url2: record[i]["company_url2"],
                            home_status: record[i]["company_status"],
                            home_status_css: record[i]["company_status_css"],
                            home_create: record[i]["company_create"],
                            home_update: record[i]["company_update"],
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
            for (var i = 0; i < record_listcurr.length; i++) {
                listCurrency = [
                    ...listCurrency,
                    {
                        curr_id: record_listcurr[i]["curr_id"],
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
    {listCurrency}
    {totalrecord}/>
{/if}