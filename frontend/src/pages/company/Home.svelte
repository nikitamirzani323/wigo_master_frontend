<script> 
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
    import Input_custom from '../../components/InputCustom.svelte' 
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let listCurrency = []
    export let listPage = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "COMPANY"
    let sData = "";
    let sDataadmin = "";
    let sDataadminrule = "";
    let myModal_newentry = "";
    let myModal_admin = "";
    let flag_btnsave = true;
    let flag_id_field = false;
    let flag_idcurr_field = false;
    let idcurr_field = "";
    let name_field = "";
    let owner_field = "";
    let email_field = "";
    let phone1_field = "";
    let phone2_field = "";
    let url1_field = "";
    let url2_field = "";
    let minimalfee_field = 0;
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = "";

    //ADMIN
    let flat_panel_active = "admin";
    let flat_panel_admin = true;
    let flat_panel_rule = false;
    let css_panel_admin = "active";
    let css_panel_rule = "";
    let listadmin = [];
    let listadminrule = [];
    let idadmincompany = 0;
    let idadminrulecompany = 0;
    let idcompany = "";
    let flag_admin_username_field = false;
    let admin_rule_field = 0;
    let admin_username_field = "";
    let admin_password_field = "";
    let admin_name_field = "";
    let admin_status_field = "";
    let admin_create_field = "";
    let admin_update_field = "";

    let adminrule_name_field = ""
    let adminrule_rule_field = ""
    let adminrule_create_field = ""
    let adminrule_update_field = ""

    let listmoney = [];
    let money_credit_field = 0;
   
    let conf_2D30_time_field = 0;
    let conf_2D30_digit_field = 0;
    let conf_2D30_minbet_field = 0;
    let conf_2D30_maxbet_field = 0;
    let conf_2D30_win_field = 0;
    let conf_2D30_maintenance_field = "";
    let conf_2D30_status_field = "";
    let conf_create_field = "";
    let conf_update_field = "";

    let pagingnow = 0;
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,idcurr,name,owner,email,phone1,phone2,minfee,url1,url2,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_id_field = true;
            flag_idcurr_field = false;
            idrecord = id;
            idcurr_field = idcurr;
            name_field = name;
            owner_field = owner;
            email_field = email;
            phone1_field = phone1;
            phone2_field = phone2;
            url1_field = url1;
            url2_field = url2;
            minimalfee_field = parseFloat(minfee);
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Currency is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(url1_field == ""){
                flag = false
                msg += "The URL 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Currency is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(url1_field == ""){
                flag = false
                msg += "The URL 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companysave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"COMPANY-SAVE",
                    company_search: searchHome,
                    company_page: parseInt(pagingnow),
                    company_id: idrecord.toUpperCase(),
                    company_idcurr: idcurr_field,
                    company_name: name_field,
                    company_owner: owner_field,
                    company_phone1: phone1_field,
                    company_phone2: phone2_field,
                    company_email: email_field,
                    company_minfee: parseFloat(minimalfee_field),
                    company_url1: url1_field,
                    company_url2: url2_field,
                    company_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function clearField(){
        flag_id_field = false;
        idrecord = "";
        idcurr_field = "";
        name_field = "";
        owner_field = "";
        email_field = "";
        phone1_field = "";
        phone2_field = "";
        url1_field = "";
        url2_field = "";
        minimalfee_field = 0;
        status_field = "";
        create_field = "";
        update_field = "";
    }
    function clearField_admin(){
        flag_admin_username_field = false;
        admin_username_field = "";
        admin_password_field = "";
        admin_name_field = "";
        admin_status_field = "";
        admin_create_field = "";
        admin_update_field = "";
    }
    function clearField_adminrule(){
        adminrule_name_field = ""
        adminrule_rule_field = ""
        adminrule_create_field = ""
        adminrule_update_field = ""
    }
    function clearField_conf(){
        conf_2D30_time_field = 0;
        conf_2D30_digit_field = 0;
        conf_2D30_minbet_field = 0;
        conf_2D30_maxbet_field = 0;
        conf_2D30_win_field = 0;
        conf_2D30_status_field = "";
        conf_create_field = "";
        conf_update_field = "";
    }

    //ADMIN
    const tab_admin = (e) => {
        if(e == "admin"){
            flat_panel_active = "admin"
            css_panel_admin = "active";
            css_panel_rule = "";
            flat_panel_admin = true;
            flat_panel_rule = false;
            call_admin(idcompany)
        }else{
            flat_panel_active = "rule"
            css_panel_admin = "";
            css_panel_rule = "active";
            flat_panel_admin = false;
            flat_panel_rule = true;
            call_adminrule(idcompany)
        }
    };
    const showAdmin = (e) => {
        idcompany = e
        call_admin(idcompany)
        myModal_admin = new bootstrap.Modal(document.getElementById("modal_companyadmin"));
        myModal_admin.show();
    };
    const call_formadmin = (e,id,idrule,username,name,status,create,update) => {
        if(flat_panel_active == "admin"){
            sDataadmin = e
            if(sDataadmin == "Edit"){
                idadmincompany = id
                admin_rule_field = idrule
                flag_admin_username_field = true;
                admin_username_field = username;
                admin_name_field = name;
                admin_status_field = status;
                admin_create_field = create;
                admin_update_field = update;
            }else{
                clearField_admin()
            }

            myModal_newentry = new bootstrap.Modal(document.getElementById("modal_crudcompanyadmin"));
            myModal_newentry.show();
        }else{
            sDataadminrule = e
            if(sDataadminrule == "Edit"){
                idadminrulecompany = id
                adminrule_name_field = name
                adminrule_rule_field = idrule
                adminrule_create_field = create
                adminrule_update_field = update
            }else{
                clearField_adminrule()
            }

            myModal_newentry = new bootstrap.Modal(document.getElementById("modal_crudcompanyadminrule"));
            myModal_newentry.show();
        }
        
        
    };
    async function call_admin(e) {
        listadmin = [];
        listadminrule = [];
        const res = await fetch("/api/companyadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                companyadmin_idcompany: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            let record_listrule = json.listrule;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listadmin = [
                        ...listadmin,
                        {
                            companyadmin_no: no,
                            companyadmin_id: record[i]["companyadmin_id"],
                            companyadmin_idrule: record[i]["companyadmin_idrule"],
                            companyadmin_idcompany: record[i]["companyadmin_idcompany"],
                            companyadmin_nmrule: record[i]["companyadmin_nmrule"],
                            companyadmin_username: record[i]["companyadmin_username"],
                            companyadmin_name: record[i]["companyadmin_name"],
                            companyadmin_status: record[i]["companyadmin_status"],
                            companyadmin_status_css: record[i]["companyadmin_status_css"],
                            companyadmin_create: record[i]["companyadmin_create"],
                            companyadmin_update: record[i]["companyadmin_update"],
                        },
                    ];
                }
            }
            for (var i = 0; i < record_listrule.length; i++) {
                listadminrule = [
                    ...listadminrule,
                    {
                        companyadminrule_id: record_listrule[i]["companyadminrule_id"],
                        companyadminrule_name: record_listrule[i]["companyadminrule_name"],
                    },
                ];
            }
        }
    }
    async function call_adminrule(e) {
        listadminrule = [];
        const res = await fetch("/api/companyadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                companyadmin_idcompany: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listadminrule = [
                        ...listadminrule,
                        {
                            companyadminrule_no: no,
                            companyadminrule_id: record[i]["companyadminrule_id"],
                            companyadminrule_nmruleadmin: record[i]["companyadminrule_nmruleadmin"],
                            companyadminrule_ruleadmin: record[i]["companyadminrule_ruleadmin"],
                            companyadminrule_create: record[i]["companyadminrule_create"],
                            companyadminrule_update: record[i]["companyadminrule_update"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave_admin() {
        let flag = true
        let msg = ""
        if(sDataadmin == "New"){
            if(idcompany == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(admin_rule_field == "" || admin_rule_field == 0){
                flag = false
                msg += "The Rule is required\n"
            }
            if(admin_username_field == ""){
                flag = false
                msg += "The Userame is required\n"
            }
            if(admin_password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
            if(admin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(admin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(admin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(admin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            let username_send = idcompany.toLowerCase()+admin_username_field.toLowerCase()
            const res = await fetch("/api/companyadminsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataadmin,
                    page:"COMPANY-SAVE",
                    companyadmin_id: parseInt(idadmincompany),
                    companyadmin_idcompany: idcompany,
                    companyadmin_idrule: parseInt(admin_rule_field),
                    companyadmin_username: username_send,
                    companyadmin_password: admin_password_field,
                    companyadmin_name: admin_name_field,
                    companyadmin_status: admin_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataadmin=="New"){
                    clearField_admin()
                }
                msgloader = json.message;
                call_admin(idcompany)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSave_adminrule() {
        let flag = true
        let msg = ""
        if(sDataadminrule == "New"){
            if(idcompany == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(adminrule_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
        }else{
            if(adminrule_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companyadminrulesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataadminrule,
                    page:"COMPANYADMIN-SAVE",
                    companyadminrule_id: parseInt(idadminrulecompany),
                    companyadminrule_idcompany: idcompany,
                    companyadminrule_nmruleadmin: adminrule_name_field,
                    companyadminrule_ruleadmin: adminrule_rule_field.toString(),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataadminrule=="New"){
                    clearField_adminrule()
                }
                msgloader = json.message;
                call_adminrule(idcompany)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    //ENDADMIN


    //MONEY
    const showMoney = (e) => {
        idcompany = e
        call_money(idcompany)
        myModal_admin = new bootstrap.Modal(document.getElementById("modal_companymoney"));
        myModal_admin.show();
    };
    async function call_money(e) {
        listmoney = [];
        const res = await fetch("/api/companymoney", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                companyadmin_idcompany: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listmoney = [
                        ...listmoney,
                        {
                            companymoney_no: no,
                            companymoney_id: record[i]["companymoney_id"],
                            companymoney_money: record[i]["companymoney_money"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave_money() {
        let flag = true
        let msg = ""
        if(idcompany == ""){
            flag = false
            msg += "The Company is required\n"
        }
        if(parseInt(money_credit_field) < 0){
            flag = false
            msg += "The Credit is required\n"
        }
       
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companymoneysave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"COMPANY-SAVE",
                    companymoney_idcompany: idcompany,
                    companymoney_money: parseInt(money_credit_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                money_credit_field = 0
                msgloader = json.message;
                call_money(idcompany)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleDelete_money(e) {
        let flag = true
        let msg = ""
        if(idcompany == ""){
            flag = false
            msg += "The Company is required\n"
        }
       
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companymoneydelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"COMPANY-SAVE",
                    companymoney_id: parseInt(e),
                    companymoney_idcompany: idcompany,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                msgloader = json.message;
                call_money(idcompany)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    //ENDMONEY

    //CONF
    const showConf = (e) => {
        clearField_conf()
        idcompany = e
        call_conf(idcompany)
        myModal_admin = new bootstrap.Modal(document.getElementById("modal_companyconf"));
        myModal_admin.show();
    };
    async function call_conf(e) {
        const res = await fetch("/api/companyconf", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                companyadmin_idcompany: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    conf_2D30_time_field = parseInt(record[i]["companyconf_2digit_30_time"]);
                    conf_2D30_digit_field = parseInt(record[i]["companyconf_2digit_30_digit"]);
                    conf_2D30_minbet_field = parseInt(record[i]["companyconf_2digit_30_minbet"]);
                    conf_2D30_maxbet_field = parseInt(record[i]["companyconf_2digit_30_maxbet"]);
                    conf_2D30_win_field = parseInt(record[i]["companyconf_2digit_30_win"]);
                    conf_2D30_maintenance_field = record[i]["companyconf_2digit_30_maintenance"];
                    conf_2D30_status_field = record[i]["companyconf_2digit_30_status"];
                    conf_create_field = record[i]["companyconf_create"];
                    conf_update_field = record[i]["companyconf_update"];
                   
                }
            }
        }
    }
    async function handleSave_conf() {
        let flag = true
        let msg = ""
        if(idcompany == ""){
            flag = false
            msg += "The Company is required\n"
        }
        if(parseInt(conf_2D30_time_field) < 0){
            flag = false
            msg += "The Time is required\n"
        }
        if(parseInt(conf_2D30_digit_field) < 0){
            flag = false
            msg += "The Digit is required\n"
        }
        if(parseInt(conf_2D30_minbet_field) < 0){
            flag = false
            msg += "The Minbet is required\n"
        }
        if(parseInt(conf_2D30_maxbet_field) < 0){
            flag = false
            msg += "The Maxbet is required\n"
        }
        if(parseFloat(conf_2D30_win_field) < 0){
            flag = false
            msg += "The Win is required\n"
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companyconfsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"COMPANY-SAVE",
                    companyconf_id: idcompany,
                    companyconf_2digit_30_time: parseInt(conf_2D30_time_field),
                    companyconf_2digit_30_digit: parseInt(conf_2D30_digit_field),
                    companyconf_2digit_30_minbet: parseInt(conf_2D30_minbet_field),
                    companyconf_2digit_30_maxbet: parseInt(conf_2D30_maxbet_field),
                    companyconf_2digit_30_win: parseFloat(conf_2D30_win_field),
                    companyconf_2digit_30_maintenance: conf_2D30_maintenance_field,
                    companyconf_2digit_30_status: conf_2D30_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                
                msgloader = json.message;
                call_conf(idcompany)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    //ENDCONF

    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterHome = [];
                listHome = [];
                const searchdata = {
                  searchHome,
                };
                dispatch("handleSearch", searchdata);
        }  
    };
    const handleSelectPaging = (event) => {
      let page = event.target.value;
      pagingnow = page;
      const pattern = {
        page
      };
      dispatch("handlePaging", pattern);
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="<i class='bi bi-arrow-clockwise'></i>&nbsp;Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-title">
                    <div class="float-end">
                      <select
                        on:change={handleSelectPaging}
                        style="text-align: center;"
                        class="form-control">
                        {#each listPage as rec}
                          <option value={rec.page_value}>{rec.page_display}</option>
                        {/each}
                      </select>
                    </div>
                  </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Code, Company"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=4>&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">STATUS</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">COMPANY</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">OWNER</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            //e,id,idcurr,name,owner,email,phone1,phone2,minfee,url1,url2,status,create,update
                                                NewData("Edit",rec.home_id,rec.home_idcurr, rec.home_name,
                                                rec.home_owner,rec.home_email,rec.home_hp1,rec.home_hp2,rec.home_minfee,
                                                rec.home_url1,rec.home_url2,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                showAdmin(rec.home_id);
                                            }} class="bi bi-person"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                showMoney(rec.home_id);
                                            }} class="bi bi-coin"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                showConf(rec.home_id);
                                            }} class="bi bi-gear"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_owner}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="currency" id="currency" 
                        disabled={flag_idcurr_field} 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listCurrency as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <Input_custom
                        bind:value={idrecord}
                        input_tipe="text_uppercase_trim"
                        input_required="required"
                        input_maxlength="5"
                        disabled={flag_id_field}
                        input_placeholder="CODE"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input_custom
                        bind:value={owner_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input_custom
                        bind:value={email_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="350"
                        input_placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone1</label>
                    <Input_custom
                        bind:value={phone1_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="20"
                        input_placeholder="Phone1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone2</label>
                    <Input_custom
                        bind:value={phone2_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="20"
                        input_placeholder="Phone2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url 1</label>
                    <Input_custom
                        bind:value={url1_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="250"
                        input_placeholder="Url1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url 2</label>
                    <Input_custom
                        bind:value={url2_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="250"
                        input_placeholder="url2"/>
                </div>
                <div class="mb-0">
                    <label for="exampleForm" class="form-label">Minimal Fee</label>
                    <Input_custom
                        bind:value={minimalfee_field}
                        input_tipe="number_float"
                        input_required=""
                        input_maxlength="5"
                        input_placeholder="Minimal Fee"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="">--Please Select--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {create_field}<br />
                        Update : {update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
        
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
    modal_id="modal_companyadmin"
    modal_size="modal-dialog-centered modal-lg"
    modal_title="ADMIN"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
    modal_footer={true}>
    <slot:template slot="search">
        <ul class="nav nav-pills">
            <li class="nav-item">
              <span on:click={() => {
                    tab_admin("admin");
                }} class="nav-link {css_panel_admin}" style="cursor: pointer;" >Admin</span>
            </li>
            <li class="nav-item">
              <span on:click={() => {
                    tab_admin("rule");
                }}  class="nav-link {css_panel_rule}" style="cursor: pointer;">Rule</span>
            </li>
        </ul>
    </slot:template>
    <slot:template slot="body">
        {#if flat_panel_admin}
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
                    <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">RULE</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                </tr>
            </thead>
            <tbody>
                {#each listadmin as rec}
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i on:click={() => {
                            //e,id,idrule,username,name,status,create,update
                                call_formadmin("Edit",rec.companyadmin_id,rec.companyadmin_idrule,rec.companyadmin_username,
                                rec.companyadmin_name,rec.companyadmin_status,rec.companyadmin_create,rec.companyadmin_update);
                            }} class="bi bi-pencil"></i>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.companyadmin_status_css}">
                            {status(rec.companyadmin_status)}
                        </span>
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_username}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_nmrule}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_name}</td>
                </tr>
                {/each}
            </tbody>
        </table>
        {/if}
        {#if flat_panel_rule}
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">RULE</th>
                </tr>
            </thead>
            <tbody>
                {#each listadminrule as rec}
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i on:click={() => {
                            //e,id,idrule,username,name,status,create,update
                                call_formadmin("Edit",rec.companyadminrule_id,rec.companyadminrule_ruleadmin,
                                "",rec.companyadminrule_nmruleadmin,"",rec.companyadminrule_create,rec.companyadminrule_update);
                            }} class="bi bi-pencil"></i>
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadminrule_nmruleadmin}</td>
                </tr>
                {/each}
            </tbody>
        </table>
        {/if}
    </slot:template>
    <slot:template slot="footer">
        <Button on:click={() => {
            call_formadmin("New");
        }} 
        button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
        button_css="btn-info"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modal_crudcompanyadmin"
	modal_size="modal-dialog-centered "
	modal_title="{sDataadmin} Admin"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Rule</label>
            <select
                bind:value="{admin_rule_field}" 
                name="adminrule" id="adminrule" 
                class="required form-control ">
                <option value="">--Please Select--</option>
                {#each listadminrule as rec}
                <option value="{rec.companyadminrule_id}">{rec.companyadminrule_name}</option>
                {/each}
            </select>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Username</label>
            <div class="input-group">
                {#if sDataadmin == "New"}
                <span class="input-group-text" id="basic-addon3">{idcompany.toLowerCase()}</span>
                {/if}
                <Input_custom
                    bind:value={admin_username_field}
                    input_tipe="text_lowercase_trim"
                    input_required="required"
                    input_maxlength="20"
                    disabled={flag_admin_username_field}
                    input_placeholder="Username"/>
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Password</label>
            <input
                bind:value={admin_password_field}
                type="password"
                class="form-control "
                placeholder="Password"
                aria-label="Password"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <Input_custom
                bind:value={admin_name_field}
                input_tipe="text_standart"
                input_required="required"
                input_maxlength="50"
                input_placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select
                class="form-control required"
                bind:value={admin_status_field}>
                <option value="">--Please Select--</option>
                <option value="Y">ACTIVE</option>
                <option value="N">DEACTIVE</option>
            </select>
        </div>
        {#if sDataadmin != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {admin_create_field}<br />
                    Update : {admin_update_field}
                </div>
            </div>
        {/if}
       
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_admin();
            }} 
            
            button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>
<Modal
	modal_id="modal_crudcompanyadminrule"
	modal_size="modal-dialog-centered "
	modal_title="{sDataadminrule} Admin Rule"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Rule</label>
            <Input_custom
                bind:value={adminrule_name_field}
                input_tipe="text_standart"
                input_required="required"
                input_maxlength="50"
                input_placeholder="Rule"/>
        </div>
        <div class="mb-3">
            <table class="table table-sm">
                <thead>
                    <tr>
                        <th colspan="2">ADMIN</th>
                        <th colspan="2">CONFIG</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td width="1%">
                            <input bind:group={adminrule_rule_field}
                                type="checkbox"
                                value="COMPANYADMIN-VIEW"/>
                        </td>
                        <td width="*">VIEW</td>
                        <td width="1%">
                            <input bind:group={adminrule_rule_field}
                                type="checkbox"
                                value="COMPANYCONFIG-VIEW"/>
                        </td>
                        <td width="*">VIEW</td>
                    </tr>
                    <tr>
                        <td width="1%">
                            <input bind:group={adminrule_rule_field}
                                type="checkbox"
                                value="COMPANYADMIN-SAVE"/>
                        </td>
                        <td width="*">SAVE</td>
                        <td width="1%">
                            <input bind:group={adminrule_rule_field}
                                type="checkbox"
                                value="COMPANYCONFIG-SAVE"/>
                        </td>
                        <td width="*">SAVE</td>
                        
                    </tr>
                </tbody>
            </table>
        </div>
        
        {#if sDataadminrule != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {adminrule_create_field}<br />
                    Update : {adminrule_update_field}
                </div>
            </div>
        {/if}
       
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
            <Button on:click={() => {
                    handleSave_adminrule();
                }} 
            button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
    modal_id="modal_companyconf"
    modal_size="modal-dialog-centered modal-lg"
    modal_title="CONFIGURE"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_footer={true}>
    <slot:template slot="body">
        <div class="accordion" id="accordionExample">
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                        Configure 2D 30S
                    </button>
                </h2>
                <div id="collapseOne" class="accordion-collapse collapse show" data-bs-parent="#accordionExample">
                    <div class="accordion-body">
                        <div class="row">
                            <div class="col-md-4">
                                <div class="mb-0">
                                    <label for="exampleForm" class="form-label">Time (s)</label>
                                    <Input_custom
                                        bind:value={conf_2D30_time_field}
                                        input_tipe="number_standart"
                                        input_required="required"
                                        input_maxlength="3"
                                        input_placeholder="Time"/>
                                </div>
                                <div class="mb-0">
                                    <label for="exampleForm" class="form-label">Digit</label>
                                    <Input_custom
                                        bind:value={conf_2D30_digit_field}
                                        input_tipe="number_standart"
                                        input_required="required"
                                        input_maxlength="4"
                                        input_placeholder="Digit"/>
                                </div>
                            </div>
                            <div class="col-md-4">
                                <div class="mb-0">
                                    <label for="exampleForm" class="form-label">MinBet</label>
                                    <Input_custom
                                        bind:value={conf_2D30_minbet_field}
                                        input_tipe="number_float"
                                        input_required="required"
                                        input_maxlength="10"
                                        input_placeholder="MinBet"/>
                                </div>
                                <div class="mb-0">
                                    <label for="exampleForm" class="form-label">MaxBet</label>
                                    <Input_custom
                                        bind:value={conf_2D30_maxbet_field}
                                        input_tipe="number_float"
                                        input_required="required"
                                        input_maxlength="510"
                                        input_placeholder="MaxBet"/>
                                </div>
                            </div>
                            <div class="col-md-4">
                                <div class="mb-0">
                                    <label for="exampleForm" class="form-label">Win</label>
                                    <Input_custom
                                        bind:value={conf_2D30_win_field}
                                        input_tipe="number_float"
                                        input_required="required"
                                        input_maxlength="5"
                                        input_placeholder="Minimal Fee"/>
                                </div>
                                <div class="mb-3">
                                    <label for="exampleForm" class="form-label">Maintenance</label>
                                    <select
                                        class="form-control required"
                                        bind:value={conf_2D30_maintenance_field}>
                                        <option value="">--Please Select--</option>
                                        <option value="Y">ACTIVE</option>
                                        <option value="N">DEACTIVE</option>
                                    </select>
                                </div>
                                <div class="mb-3">
                                    <label for="exampleForm" class="form-label">Status</label>
                                    <select
                                        class="form-control required"
                                        bind:value={conf_2D30_status_field}>
                                        <option value="">--Please Select--</option>
                                        <option value="Y">ACTIVE</option>
                                        <option value="N">DEACTIVE</option>
                                    </select>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
                        Accordion Item #2
                    </button>
                </h2>
                <div id="collapseTwo" class="accordion-collapse collapse" data-bs-parent="#accordionExample">
                    <div class="accordion-body">
                        <strong>This is the second item's accordion body.</strong> It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions. You can modify any of this with custom CSS or overriding our default variables. It's also worth noting that just about any HTML can go within the <code>.accordion-body</code>, though the transition does limit overflow.
                    </div>
                </div>
            </div>
        </div>
        <div class="mb-3 mt-4">
            <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                Create : {conf_create_field}<br />
                Update : {conf_update_field}
            </div>
        </div>
    </slot:template>
    <slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_conf();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
    modal_id="modal_companymoney"
    modal_size="modal-dialog-centered "
    modal_title="CREDIT"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_footer={false}>
    <slot:template slot="body">
        <div class="row g-2">
            <div class="col-auto">
                <label for="exampleForm" class="visually-hidden">Credit</label>
                <Input_custom
                    bind:value={money_credit_field}
                    input_tipe="number_float"
                    input_required="required"
                    input_maxlength="10"
                    input_placeholder="MinBet"/>
            </div>
            <div class="col-auto ">
                <Button on:click={() => {
                        handleSave_money();
                    }} 
                    button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
                    button_css="btn-warning mt-1"/>
            </div>
        </div>
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CREDIT</th>
                </tr>
            </thead>
            <tbody>
                {#each listmoney as rec}
                <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                        <i on:click={() => {
                            handleDelete_money(rec.companymoney_id,);
                            }} class="bi bi-trash3"></i>
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.companymoney_money)}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
    <slot:template slot="footer">
        
	</slot:template>
</Modal>
