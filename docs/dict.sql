/*==============================================================*/
/* Table: dict_data                                         */
/*==============================================================*/
create table dict_data
(
    id         int(11) not null auto_increment comment '字典编码',
    sort       int(4)       default 0 comment '字典排序',
    label      varchar(100) default '' comment '字典标签',
    value      varchar(100) default '' comment '字典键值',
    type       varchar(100) default '' comment '字典类型',
    css_class  varchar(100) default '' comment '样式属性（其他样式扩展）',
    list_class varchar(100) default '' comment '表格回显样式',
    is_default char(1)      default 'N' comment '是否默认（Y是 N否）',
    status     char(1)      default '0' comment '状态（1正常 2停用）',
    create_at  datetime     default NULL comment '创建时间',
    create_by  varchar(64)  default '' comment '创建者',
    update_at  datetime     default NULL comment '更新时间',
    update_by  varchar(64)  default '' comment '更新者',
    remark     varchar(500) default '' comment '备注',
    primary key (id)
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 410
    DEFAULT CHARSET = utf8
    ROW_FORMAT = COMPACT COMMENT ='字典数据表';

alter table dict_data
    comment '字典数据表';

/*==============================================================*/
/* Table: dict_type                                         */
/*==============================================================*/
create table dict_type
(
    dict_id     bigint(20) not null auto_increment comment '字典主键',
    dict_name   varchar(100) default '' comment '字典名称',
    dict_type   varchar(100) default '' comment '字典类型',
    status      char(1)      default '0' comment '状态（1正常 2停用）',
    create_by   varchar(64)  default '' comment '创建者',
    create_time datetime     default NULL comment '创建时间',
    update_by   varchar(64)  default '' comment '更新者',
    update_time datetime     default NULL comment '更新时间',
    remark      varchar(500) default NULL comment '备注',
    primary key (dict_id),
    unique key dict_type (dict_type)
)
    ENGINE = InnoDB
    AUTO_INCREMENT = 131
    DEFAULT CHARSET = utf8 COMMENT ='字典类型表';

alter table dict_type
    comment '字典类型表';
