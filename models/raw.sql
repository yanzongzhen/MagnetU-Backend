-- 创建文件仓库表
CREATE TABLE Repositories (
                              RepositoryID INT PRIMARY KEY,
                              UserID INT,
                              CurrentCapacity BIGINT,
                              MaxCapacity BIGINT,
                              Permissions VARCHAR(255)
    -- 可能还需要添加对用户表的外键约束
    -- FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- 创建文件夹表
CREATE TABLE Folders (
                         FolderID INT PRIMARY KEY,
                         FolderName VARCHAR(255),
                         ParentFolderID INT,
                         RepositoryID INT,
                         CreationTime DATETIME,
                         FOREIGN KEY (ParentFolderID) REFERENCES Folders(FolderID),
                         FOREIGN KEY (RepositoryID) REFERENCES Repositories(RepositoryID)
);

-- 创建文件表
CREATE TABLE Files (
                       FileID INT PRIMARY KEY,
                       FileName VARCHAR(255),
                       ParentFolderID INT,
                       RepositoryID INT,
                       FileExtension VARCHAR(255),
                       FileType VARCHAR(255),
                       FileSize BIGINT,
                       UploadTime DATETIME,
                       DownloadCount INT,
                       StoragePath VARCHAR(255),
                       FOREIGN KEY (ParentFolderID) REFERENCES Folders(FolderID),
                       FOREIGN KEY (RepositoryID) REFERENCES Repositories(RepositoryID)
);
