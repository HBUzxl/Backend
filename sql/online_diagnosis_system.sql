/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80403
 Source Host           : localhost:3306
 Source Schema         : online_diagnosis_system

 Target Server Type    : MySQL
 Target Server Version : 80403
 File Encoding         : 65001

 Date: 18/01/2025 10:38:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nick_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_admins_id`(`id`) USING BTREE,
  UNIQUE INDEX `uni_admins_username`(`username`) USING BTREE,
  INDEX `idx_admins_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------

-- ----------------------------
-- Table structure for allocators
-- ----------------------------
DROP TABLE IF EXISTS `allocators`;
CREATE TABLE `allocators`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nick_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_allocators_id`(`id`) USING BTREE,
  UNIQUE INDEX `uni_allocators_username`(`username`) USING BTREE,
  INDEX `idx_allocators_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of allocators
-- ----------------------------
INSERT INTO `allocators` VALUES (1, 'admin', 'admin1', NULL, '分配员', NULL, '2024-12-24 14:54:20.059', NULL);

-- ----------------------------
-- Table structure for appointments
-- ----------------------------
DROP TABLE IF EXISTS `appointments`;
CREATE TABLE `appointments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `appointment_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `patient_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_gender` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_age` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `appointment_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `surgery_location` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `clinical_doctor` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `expert_id` bigint UNSIGNED NULL DEFAULT NULL,
  `hospital` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `remarks` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `submit_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `appointment_status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_appointments_expert_id`(`expert_id`) USING BTREE,
  INDEX `idx_appointments_deleted_at`(`deleted_at`) USING BTREE,
  CONSTRAINT `fk_appointments_expert` FOREIGN KEY (`expert_id`) REFERENCES `experts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of appointments
-- ----------------------------
INSERT INTO `appointments` VALUES (1, 'AP001', '张三11', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 101, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 13:34:05', '申请中', NULL);
INSERT INTO `appointments` VALUES (2, 'AP002', '李四', '女', '42', '9876543210', '2023-10-16 11:30:00', '心脏', '王医生', 102, '北京协和医院', '需进行心脏手术', '2024-12-24 15:10:30', '申请中', NULL);
INSERT INTO `appointments` VALUES (3, 'AP003', '王五', '男', '50', '5556667777', '2023-10-17 09:00:00', '肺', '刘医生', 103, '上海中山医院', '需进一步检查肺部CT', '2024-12-25 16:45:22', '申请中', NULL);
INSERT INTO `appointments` VALUES (4, 'AP004', '赵六', '女', '30', '1112223333', '2023-10-18 14:00:00', '肝', '陈医生', 104, '广州医科大学附属第一医院', '需进行肝功能检查', '2024-12-26 17:30:15', '已预约', NULL);
INSERT INTO `appointments` VALUES (5, 'AP005', '孙七', '男', '60', '4445556666', '2023-10-19 16:00:00', '肾', '李医生', 105, '四川大学华西医院', '需进行肾功能评估', '2024-12-27 18:20:40', '已预约', NULL);
INSERT INTO `appointments` VALUES (19, 'AP_20241224112903', '张三1', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 102, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 11:29:03', '申请中', NULL);
INSERT INTO `appointments` VALUES (20, 'AP_20241224113116', '张三3', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 101, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 11:31:17', '申请中', NULL);
INSERT INTO `appointments` VALUES (21, 'AP_20241224113352', '张三4', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 102, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 11:33:52', '申请中', NULL);
INSERT INTO `appointments` VALUES (22, 'AP_20241224113817', '张三', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 101, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 11:38:17', '申请中', NULL);
INSERT INTO `appointments` VALUES (23, 'AP_20241224114350', '张三6', '男', '35', '1234567890', '2023-10-15 10:00:00', '胃', '张医生', 101, '华东病理远程诊断中心', '需进一步检查', '2024-12-24 11:43:50', '申请中', NULL);

-- ----------------------------
-- Table structure for attachments
-- ----------------------------
DROP TABLE IF EXISTS `attachments`;
CREATE TABLE `attachments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `file_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `file_path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `file_size` bigint NULL DEFAULT NULL,
  `case_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `attachment_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_attachments_id`(`id`) USING BTREE,
  INDEX `fk_cases_attachments`(`case_id`) USING BTREE,
  CONSTRAINT `fk_cases_attachments` FOREIGN KEY (`case_id`) REFERENCES `cases` (`case_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of attachments
-- ----------------------------
INSERT INTO `attachments` VALUES (1, 'fig2.jpg', 'uploads\\attachments\\case_C001\\fig2.jpg', 29530, 'C001', '');

-- ----------------------------
-- Table structure for cases
-- ----------------------------
DROP TABLE IF EXISTS `cases`;
CREATE TABLE `cases`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `case_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `patient_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_gender` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_age` bigint NULL DEFAULT NULL,
  `patient_phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `biopsy_site` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `tissue_count` bigint NULL DEFAULT NULL,
  `bar_code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `checkup_no` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `clinical_phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `hospital` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `sample_date` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `receive_date` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `pathology_type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `inpatient_no` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `bed_no` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `marital_status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `patient_address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `clinical_diagnosis` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `clinical_data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `gross_finding` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `immunohistochemistry` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `pathological_diagnosis` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `remarks` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `print_count` bigint NULL DEFAULT NULL,
  `case_status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `expert_id` bigint UNSIGNED NULL DEFAULT NULL,
  `consultation_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `submit_at` datetime NULL DEFAULT NULL,
  `diagnose_at` datetime NULL DEFAULT NULL,
  `diagnosis_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `expert_diagnosis_opinion` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `mirror_description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `diagnosis_remarks` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_cases_id`(`id`) USING BTREE,
  UNIQUE INDEX `uni_cases_case_id`(`case_id`) USING BTREE,
  INDEX `idx_cases_expert_id`(`expert_id`) USING BTREE,
  INDEX `idx_cases_deleted_at`(`deleted_at`) USING BTREE,
  CONSTRAINT `fk_experts_cases` FOREIGN KEY (`expert_id`) REFERENCES `experts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cases
-- ----------------------------
INSERT INTO `cases` VALUES (1, 'C001', '张三1', '男', 45, '13812345678', '门诊', '肺部', 2, 'BAR001', 'CHK001', '13812345678', '华东病理远程诊断中心', '2023-09-01', '2023-09-03', '肺癌', 'IP001', '101', '已婚', '北京市朝阳区', '肺部肿块', 'CT扫描显示肺部阴影', '肿块大小3cm', '阳性', '肺癌', '无', 5, 'pendingdiagnosis', 101, 'HZ_C001', '2025-01-10 10:32:03', '2025-01-10 18:27:17', '阴性', NULL, '1123', '', '');
INSERT INTO `cases` VALUES (2, 'C002', '李四', '女', 32, '13987654321', '住院', '乳腺', 1, 'BAR002', 'CHK002', '13987654321', '华东病理远程诊断中心', '2023-09-02', '2023-09-04', '乳腺癌', 'IP002', '202', '未婚', '上海市浦东新区', '乳腺结节', '超声显示乳腺肿块', '肿块大小2cm', '阴性', '乳腺纤维瘤', '需进一步检查', 3, 'pendingdiagnosis', 101, 'HZ_C002', '2025-01-10 10:31:33', '2024-12-22 16:09:42', '阴性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (3, 'C003', '王五1', '男', 50, '13712345678', '门诊', '胃部', 3, 'BAR003', 'CHK003', '13712345678', '华东病理远程诊断中心', '2023-09-03', '2023-09-05', '胃癌', 'IP003', '303', '已婚', '广州市天河区', '胃部不适', '胃镜检查显示胃部溃疡', '溃疡面积2cm*3cm', '阳性', '胃癌', '无', 0, 'pendingdiagnosis', 101, 'HZ_C003', '2024-12-22 17:50:07', '2025-01-10 16:44:25', '阴性', NULL, '22', '', '22');
INSERT INTO `cases` VALUES (4, 'C004', '赵六', '女', 60, '13612345678', '住院', '甲状腺', 1, 'BAR004', 'CHK004', '13612345678', '华东病理远程诊断中心', '2023-09-04', '2023-09-06', '甲状腺癌', 'IP004', '404', '已婚', '成都市武侯区', '甲状腺肿大', '超声显示甲状腺结节', '结节大小1.5cm', '阳性', '甲状腺癌', '无', 3, 'unsubmitted', 102, 'HZ_C004', '2024-12-22 17:50:13', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (5, 'C005', '孙七', '男', 40, '13512345678', '门诊', '肝脏', 2, 'BAR005', 'CHK005', '13512345678', '瑞金医院', '2023-09-05', '2023-09-07', '肝癌', 'IP005', '505', '未婚', '深圳市福田区', '肝区疼痛', 'CT扫描显示肝脏占位', '占位大小4cm', '阳性', '肝癌', '无', 1, 'unsubmitted', 105, 'HZ_C005', '2024-12-22 17:50:13', '2024-12-22 16:09:42', '阴性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (6, 'C006', '周八', '男', 55, '13412345678', '住院', '肺部', 2, 'BAR006', 'CHK006', '13412345678', '协和医院', '2023-09-06', '2023-09-08', '肺癌', 'IP006', '106', '已婚', '北京市朝阳区', '咳嗽咳痰', 'CT扫描显示肺部阴影', '阴影面积5cm*6cm', '阳性', '肺癌', '无', 2, 'unsubmitted', 106, 'HZ_C006', '2024-12-22 17:50:13', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (7, 'C007', '吴九', '女', 48, '13312345678', '住院', '乳腺', 1, 'BAR007', 'CHK007', '13312345678', '协和医院', '2023-09-07', '2023-09-09', '乳腺癌', 'IP007', '207', '未婚', '上海市静安区', '乳腺结节', '超声显示乳腺肿块', '肿块大小3cm', '阳性', '乳腺癌', '无', 4, 'pendingdiagnosis', 101, 'HZ_C007', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (8, 'C008', '郑十', '男', 62, '13212345678', '门诊', '胃部', 3, 'BAR008', 'CHK008', '13212345678', '中山医院', '2023-09-08', '2023-09-10', '胃癌', 'IP008', '308', '已婚', '广州市越秀区', '胃部不适', '胃镜检查显示胃部溃疡', '溃疡面积3cm*4cm', '阳性', '胃癌', '无', 4, 'pendingdiagnosis', 101, 'HZ_C008', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阴性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (9, 'C009', '王十一', '女', 35, '13112345678', '门诊', '甲状腺', 1, 'BAR009', 'CHK009', '13112345678', '华西医院', '2023-09-09', '2023-09-11', '甲状腺癌', 'IP009', '409', '已婚', '成都市锦江区', '甲状腺肿大', '超声显示甲状腺结节', '结节大小2cm', '阳性', '甲状腺癌', '需进一步检查', 3, 'returned', 109, 'HZ_C009', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (10, 'C010', '李十二', '男', 42, '13012345678', '住院', '肝脏', 2, 'BAR010', 'CHK010', '13012345678', '瑞金医院', '2023-09-10', '2023-09-12', '肝癌', 'IP010', '510', '未婚', '深圳市南山区', '肝区疼痛', 'CT扫描显示肝脏占位', '占位大小5cm', '阳性', '肝癌', '无', 1, 'returned', 110, 'HZ_C010', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (11, 'C011', '赵十三', '女', 58, '12912345678', '门诊', '肺部', 3, 'BAR011', 'CHK011', '12912345678', '人民医院', '2023-09-11', '2023-09-13', '肺癌', 'IP011', '111', '已婚', '北京市西城区', '肺部肿块', 'CT扫描显示肺部阴影', '肿块大小5cm', '阳性', '肺癌', '无', 5, 'pendingdiagnosis', 101, 'HZ_C011', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (12, 'C012', '钱十四', '男', 47, '12812345678', '住院', '乳腺', 1, 'BAR012', 'CHK012', '12812345678', '协和医院', '2023-09-12', '2023-09-14', '乳腺癌', 'IP012', '212', '未婚', '上海市黄浦区', '乳腺结节', '超声显示乳腺肿块', '肿块大小2.5cm', '阳性', '乳腺癌', '无', 6, 'withdraw', 101, 'HZ_C012', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (13, 'C013', '孙十五', '女', 53, '12712345678', '门诊', '胃部', 3, 'BAR013', 'CHK013', '12712345678', '中山医院', '2023-09-13', '2023-09-15', '胃癌', 'IP013', '313', '已婚', '广州市海珠区', '胃部不适', '胃镜检查显示胃部溃疡', '溃疡面积4cm*5cm', '阳性', '胃癌', '无', 1, 'withdraw', 103, 'HZ_C013', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);
INSERT INTO `cases` VALUES (14, 'C014', '周十六', '男', 65, '12612345678', '住院', '甲状腺', 1, 'BAR014', 'CHK014', '12612345678', '华西医院', '2023-09-14', '2023-09-16', '甲状腺癌', 'IP014', '414', '已婚', '成都市青羊区', '甲状腺肿大', '超声显示甲状腺结节', '结节大小2.5cm', '阳性', '甲状腺癌', '无', 23, 'withdraw', 104, 'HZ_C014', '2024-12-22 16:09:42', '2024-12-22 16:09:42', '阳性', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for experts
-- ----------------------------
DROP TABLE IF EXISTS `experts`;
CREATE TABLE `experts`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nick_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `hospital` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_experts_id`(`id`) USING BTREE,
  UNIQUE INDEX `uni_experts_username`(`username`) USING BTREE,
  INDEX `idx_experts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 110 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of experts
-- ----------------------------
INSERT INTO `experts` VALUES (101, 'expert1', '111111', 'expert', '张医生', '华东病理远程诊断中心', '13812345678', '2023-01-01 10:00:00.000', '2024-12-27 11:01:27.210', NULL);
INSERT INTO `experts` VALUES (102, 'expert2', 'securepass456', 'expert', '李医生', '华东病理远程诊断中心', '13812345678', '2023-01-02 11:00:00.000', '2023-01-02 11:00:00.000', NULL);
INSERT INTO `experts` VALUES (103, 'expert3', 'expertpass789', 'expert', '王医生', '华东病理远程诊断中心', '13812345678', '2023-01-03 12:00:00.000', '2023-01-03 12:00:00.000', NULL);
INSERT INTO `experts` VALUES (104, 'expert4', 'strongpass101', 'expert', '赵医生', '华东病理远程诊断中心', '13812345678', '2023-01-04 13:00:00.000', '2023-01-04 13:00:00.000', NULL);
INSERT INTO `experts` VALUES (105, 'expert5', 'safe12345', 'expert', '孙医生', '瑞金医院', '13812345678', '2023-01-05 14:00:00.000', '2023-01-05 14:00:00.000', NULL);
INSERT INTO `experts` VALUES (106, 'expert6', 'docpass123', 'expert', '周医生', '人民医院', '13812345678', '2023-01-06 15:00:00.000', '2023-01-06 15:00:00.000', NULL);
INSERT INTO `experts` VALUES (107, 'expert7', 'securedoc456', 'expert', '吴医生', '协和医院', '13812345678', '2023-01-07 16:00:00.000', '2023-01-07 16:00:00.000', NULL);
INSERT INTO `experts` VALUES (108, 'expert8', 'docpass789', 'expert', '郑医生', '中山医院', '13812345678', '2023-01-08 17:00:00.000', '2023-01-08 17:00:00.000', NULL);
INSERT INTO `experts` VALUES (109, 'expert9', 'strongdoc101', 'expert', '王医生', '华东病理远程诊断中心', '13812345678', '2023-01-09 18:00:00.000', '2023-01-09 18:00:00.000', NULL);
INSERT INTO `experts` VALUES (110, 'expert10', 'safedoc12345', 'expert', '李医生', '华东病理远程诊断中心', '13812345678', '2023-01-10 19:00:00.000', '2023-01-10 19:00:00.000', NULL);

-- ----------------------------
-- Table structure for slices
-- ----------------------------
DROP TABLE IF EXISTS `slices`;
CREATE TABLE `slices`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `slice_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `file_size` bigint NULL DEFAULT NULL,
  `case_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `file_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_cases_slices`(`case_id`) USING BTREE,
  CONSTRAINT `fk_cases_slices` FOREIGN KEY (`case_id`) REFERENCES `cases` (`case_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of slices
-- ----------------------------
INSERT INTO `slices` VALUES (1, 'C001_fg1.jpg', 'fg1.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 90292, 'C001', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (2, 'C001_fig2.jpg', 'fig2.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 29530, 'C001', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (3, 'C002_SE_RIS_100.jpg', 'SE_RIS_100.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg', 46335, 'C002', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg');
INSERT INTO `slices` VALUES (4, 'C005_SE_RIS_100.jpg', 'SE_RIS_100.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg', 46335, 'C005', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg');
INSERT INTO `slices` VALUES (5, 'C001_QQ截图20241128164953.png', 'QQ截图20241128164953.png', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg', 58286, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg');
INSERT INTO `slices` VALUES (6, 'C001_QQ截图20241128165508.png', 'QQ截图20241128165508.png', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg', 269830, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg');
INSERT INTO `slices` VALUES (7, 'C001_fig2.jpg', 'fig2.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 29530, 'C001', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (8, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 66396, 'C003', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (9, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg');
INSERT INTO `slices` VALUES (10, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg');
INSERT INTO `slices` VALUES (11, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg');
INSERT INTO `slices` VALUES (12, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg');
INSERT INTO `slices` VALUES (13, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg');
INSERT INTO `slices` VALUES (14, 'C003_nmse_snr.jpg', 'nmse_snr.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg', 66396, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg');
INSERT INTO `slices` VALUES (15, 'C003_QQ截图20241128165508.png', 'QQ截图20241128165508.png', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg', 269830, 'C003', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg');
INSERT INTO `slices` VALUES (16, 'C002_IMG006x013.jpg', 'IMG006x013.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 976622, 'C002', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (17, 'C002_IMG006x013.jpg', 'IMG006x013.jpg', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg', 976622, 'C002', '\\uploads\\slices\\case_C002\\C002_IMG006x013.jpg');
INSERT INTO `slices` VALUES (18, 'C001_IMG006x013.jpg', 'IMG006x013.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg', 976622, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg');
INSERT INTO `slices` VALUES (19, 'C001_IMG006x015.jpg', 'IMG006x015.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg', 991313, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg');
INSERT INTO `slices` VALUES (20, 'C001_IMG006x016.jpg', 'IMG006x016.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg', 874038, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg');
INSERT INTO `slices` VALUES (21, 'C001_IMG006x014.jpg', 'IMG006x014.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg', 946994, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg');
INSERT INTO `slices` VALUES (22, 'C001_IMG006x013.jpg', 'IMG006x013.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg', 976622, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x013.jpg');
INSERT INTO `slices` VALUES (23, 'C001_IMG006x016.jpg', 'IMG006x016.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg', 874038, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x016.jpg');
INSERT INTO `slices` VALUES (24, 'C001_IMG006x015.jpg', 'IMG006x015.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg', 991313, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x015.jpg');
INSERT INTO `slices` VALUES (25, 'C001_IMG006x014.jpg', 'IMG006x014.jpg', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg', 946994, 'C001', '\\uploads\\slices\\case_C001\\C001_IMG006x014.jpg');

SET FOREIGN_KEY_CHECKS = 1;
