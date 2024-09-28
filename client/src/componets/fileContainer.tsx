import { useQuery } from "@tanstack/react-query";
import FilesList from "./fileslist";

interface data {
  fileUrl: string;
  fileType: string;
  fileName: string;
  fileId: string;
}

const getfile = async () => {
  const response = await fetch("/api/user/getfile", {
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json() as Promise<data[]>;
};
const FileContainer = () => {
  const { isLoading, data } = useQuery({
    queryKey: ["files"],
    queryFn: getfile,
  });
  console.log(isLoading, data);
  return (
    <>
      {data &&
        data.map((data) => {
          return (
            <FilesList
              key={data.fileId}
              fileName={data.fileName}
              fileUrl={data.fileUrl}
              download={downloadFile}
              id={data.fileId}
            />
          );
        })}
    </>
  );
};

export default FileContainer;

async function downloadFile(fileurl: URL, fileName: string) {
  try {
    const response = await fetch(fileurl);
    if (!response.ok) throw new Error("Network response was not ok.");

    console.log(fileName);
    const blob = await response.blob();

    const url = window.URL.createObjectURL(blob);

    const a = document.createElement("a");
    a.href = url;
    a.download = fileName; // Specify the filename for the download
    document.body.appendChild(a);
    a.click();

    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Error downloading file:", error);
  }
}
